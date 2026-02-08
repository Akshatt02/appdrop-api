// Package services provides business logic and validation for the AppDrop API.
// Services layer sits between handlers and repository, enforcing business rules
// and complex validation logic before database operations.
package services

import (
	"appdrop-api/internal/models"
	"appdrop-api/internal/repository"
	"errors"
)

// GetPages retrieves all pages from the database.
// Returns a list of all pages or an error if database operation fails.
func GetPages() ([]models.Page, error) {
	return repository.GetAllPages()
}

func CreatePage(page models.Page) (*models.Page, error) {

	// CreatePage validates and creates a new page.
	// Business Rules Enforced:
	//   - Name and route are required (non-empty strings)
	//   - Route must be globally unique
	//   - If is_home=true, ensures only one home page by resetting others
	// Returns the created page with its UUID or an error.

	if page.Name == "" || page.Route == "" {
		return nil, errors.New("name and route are required")
	}

	exists, err := repository.RouteExists(page.Route)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("page route already exists")
	}

	if page.IsHome {
		err := repository.ResetHomePage()
		if err != nil {
			return nil, err
		}
	}

	return repository.CreatePage(page)
}

func GetPageWithWidgets(id string) (map[string]interface{}, error) {
	// GetPageWithWidgets retrieves a page and all its associated widgets.
	// Returns a map containing both page details and widgets array.
	// Ensures widgets array is empty array instead of null.
	page, err := repository.GetPageByID(id)
	if err != nil {
		return nil, err
	}

	widgets, err := repository.GetWidgetsByPageID(id)
	if err != nil {
		return nil, err
	}

	// Ensure empty array instead of null
	if widgets == nil {
		widgets = []models.Widget{}
	}

	response := map[string]interface{}{
		"page":    page,
		"widgets": widgets,
	}

	return response, nil
}

func DeletePage(id string) error {
	// DeletePage removes a page from the database.
	// Business Rule: Cannot delete the home page (is_home=true).
	// Returns error if page not found or if attempting to delete home page.
	page, err := repository.GetPageByID(id)
	if err != nil {
		return errors.New("page not found")
	}

	// Rule: cannot delete home page
	if page.IsHome {
		return errors.New("cannot delete home page")
	}

	return repository.DeletePage(id)
}

func UpdatePage(id string, page models.Page) (*models.Page, error) {

	// UpdatePage modifies an existing page's details.
	// Business Rules Enforced:
	//   - Name and route are required (non-empty strings)
	//   - Page must exist
	//   - Route must be unique (excluding the current page)
	//   - If is_home=true, ensures only one home page by resetting others
	// Returns the updated page or an error.

	if page.Name == "" || page.Route == "" {
		return nil, errors.New("name and route are required")
	}

	// check page exists
	_, err := repository.GetPageByID(id)
	if err != nil {
		return nil, errors.New("page not found")
	}

	// route must be unique (excluding same page)
	exists, err := repository.RouteExistsForOtherPage(page.Route, id)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("page route already exists")
	}

	// only one home page rule
	if page.IsHome {
		err := repository.ResetHomePage()
		if err != nil {
			return nil, err
		}
	}

	page.ID = id
	return repository.UpdatePage(page)
}
