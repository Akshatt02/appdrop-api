package handlers

import (
	"encoding/json"
	"net/http"

	"appdrop-api/internal/models"
	"appdrop-api/internal/services"
	"appdrop-api/internal/utils"
)

func CreateWidgetHandler(w http.ResponseWriter, r *http.Request) {
	pageID := r.URL.Path[len("/pages/"):]
	pageID = pageID[:len(pageID)-len("/widgets")]

	var widget models.Widget
	err := json.NewDecoder(r.Body).Decode(&widget)
	if err != nil {
		utils.SendError(w, 400, "INVALID_JSON", "Invalid request body")
		return
	}

	widget.PageID = pageID

	createdWidget, err := services.CreateWidget(widget)
	if err != nil {
		if err.Error() == "page not found" {
			utils.SendError(w, 404, "NOT_FOUND", "Page not found")
		} else {
			utils.SendError(w, 400, "VALIDATION_ERROR", err.Error())
		}
		return
	}

	utils.SendJSON(w, 201, createdWidget)
}

func UpdateWidgetHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/widgets/"):]

	var widget models.Widget
	err := json.NewDecoder(r.Body).Decode(&widget)
	if err != nil {
		utils.SendError(w, 400, "INVALID_JSON", "Invalid request body")
		return
	}
	widget.ID = id

	updatedWidget, err := services.UpdateWidget(widget)
	if err != nil {
		if err.Error() == "widget not found" {
			utils.SendError(w, 404, "NOT_FOUND", "Widget not found")
		} else {
			utils.SendError(w, 400, "VALIDATION_ERROR", err.Error())
		}
		return
	}

	utils.SendJSON(w, 200, updatedWidget)
}

func DeleteWidgetHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/widgets/"):]

	err := services.DeleteWidget(id)
	if err != nil {
		if err.Error() == "widget not found" {
			utils.SendError(w, 404, "NOT_FOUND", "Widget not found")
		} else {
			utils.SendError(w, 400, "VALIDATION_ERROR", err.Error())
		}
		return
	}

	utils.SendJSON(w, 200, map[string]string{"message": "Widget deleted"})
}

func ReorderWidgetsHandler(w http.ResponseWriter, r *http.Request) {
	pageID := r.URL.Path[len("/pages/"):]
	pageID = pageID[:len(pageID)-len("/widgets/reorder")]

	var body struct {
		WidgetIDs []string `json:"widget_ids"`
	}

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		utils.SendError(w, 400, "INVALID_JSON", "Invalid request body")
		return
	}

	err = services.ReorderWidgets(pageID, body.WidgetIDs)
	if err != nil {
		switch err.Error() {
		case "page not found":
			utils.SendError(w, 404, "NOT_FOUND", "Page not found")
		case "widget not found":
			utils.SendError(w, 404, "NOT_FOUND", "One or more widgets not found")
		default:
			utils.SendError(w, 400, "VALIDATION_ERROR", err.Error())
		}
		return
	}

	utils.SendJSON(w, 200, map[string]string{"message": "Widgets reordered"})
}
