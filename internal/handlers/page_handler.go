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

	utils.SendJSON(w, 200, pages)
}

func CreatePageHandler(w http.ResponseWriter, r *http.Request) {
	var page models.Page

	err := json.NewDecoder(r.Body).Decode(&page)
	if err != nil {
		utils.SendError(w, 400, "INVALID_JSON", "Invalid request body")
		return
	}

	err = services.CreatePage(page)
	if err != nil {
		utils.SendError(w, 400, "VALIDATION_ERROR", err.Error())
		return
	}

	utils.SendJSON(w, 201, map[string]string{"message": "Page created"})
}
