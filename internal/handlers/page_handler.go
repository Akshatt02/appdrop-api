// Package handlers provides HTTP request handlers for the AppDrop API.
// Each handler parses incoming requests, calls appropriate service methods,
// and returns formatted responses with proper HTTP status codes.
package handlers

import (
	"encoding/json"
	"net/http"

	"appdrop-api/internal/models"
	"appdrop-api/internal/services"
	"appdrop-api/internal/utils"
)

// GetPagesHandler handles GET /pages requests.
// Returns a list of all pages in the application.
// Returns an empty array if no pages exist (never null).
// Status: 200 OK on success, 500 on database error
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

// CreatePageHandler handles POST /pages requests.
// Creates a new page with provided name, route, and is_home status.
// Validates request body, route uniqueness, and is_home constraints.
// Returns the created page with its UUID.
// Status: 201 Created on success, 400 for validation errors
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

// GetPageByIDHandler handles GET /pages/:id requests.
// Retrieves a page by UUID along with all its associated widgets.
// Returns complete page structure including widget array.
// Status: 200 OK on success, 404 if page not found
func GetPageByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/pages/"):]

	data, err := services.GetPageWithWidgets(id)
	if err != nil {
		utils.SendError(w, 404, "NOT_FOUND", "Page not found")
		return
	}

	utils.SendJSON(w, 200, data)
}

// DeletePageHandler handles DELETE /pages/:id requests.
// Deletes a page and cascades delete to all its widgets.
// Cannot delete the page marked as is_home=true.
// Status: 200 OK on success, 404 if page not found, 409 if trying to delete home page
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

// UpdatePageHandler handles PUT /pages/:id requests.
// Updates page name, route, or is_home status.
// Validates new route uniqueness and is_home constraints.
// Returns the updated page.
// Status: 200 OK on success, 404 if page not found, 409 if route conflict
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
