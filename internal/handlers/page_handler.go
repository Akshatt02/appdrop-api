package handlers

import (
	"encoding/json"
	"net/http"

	"appdrop-api/internal/models"
	"appdrop-api/internal/services"
	"appdrop-api/internal/utils"
)

func GetPagesHandler(w http.ResponseWriter, r *http.Request) {
	pages, err := services.GetPages()
	if err != nil {
		utils.SendError(w, 500, "INTERNAL_ERROR", err.Error())
		return
	}

	// Ensure empty array instead of null
	if pages == nil {
		pages = []models.Page{}
	}

	utils.SendJSON(w, 200, pages)
}

func CreatePageHandler(w http.ResponseWriter, r *http.Request) {
	var page models.Page

	err := json.NewDecoder(r.Body).Decode(&page)
	if err != nil {
		utils.SendError(w, 400, "INVALID_JSON", "Invalid request body")
		return
	}

	createdPage, err := services.CreatePage(page)
	if err != nil {
		utils.SendError(w, 400, "VALIDATION_ERROR", err.Error())
		return
	}

	utils.SendJSON(w, 201, createdPage)
}

func GetPageByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/pages/"):]

	data, err := services.GetPageWithWidgets(id)
	if err != nil {
		utils.SendError(w, 404, "NOT_FOUND", "Page not found")
		return
	}

	utils.SendJSON(w, 200, data)
}

func DeletePageHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/pages/"):]

	err := services.DeletePage(id)
	if err != nil {
		switch err.Error() {
		case "cannot delete home page":
			utils.SendError(w, 409, "CONFLICT", "Cannot delete home page")
		case "page not found":
			utils.SendError(w, 404, "NOT_FOUND", "Page not found")
		default:
			utils.SendError(w, 400, "VALIDATION_ERROR", err.Error())
		}
		return
	}

	utils.SendJSON(w, 200, map[string]string{"message": "Page deleted"})
}

func UpdatePageHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/pages/"):]

	var page models.Page
	err := json.NewDecoder(r.Body).Decode(&page)
	if err != nil {
		utils.SendError(w, 400, "INVALID_JSON", "Invalid request body")
		return
	}

	updatedPage, err := services.UpdatePage(id, page)
	if err != nil {
		switch err.Error() {
		case "page not found":
			utils.SendError(w, 404, "NOT_FOUND", "Page not found")
		case "page route already exists":
			utils.SendError(w, 409, "CONFLICT", "Page route already exists")
		default:
			utils.SendError(w, 400, "VALIDATION_ERROR", err.Error())
		}
		return
	}

	utils.SendJSON(w, 200, updatedPage)
}
