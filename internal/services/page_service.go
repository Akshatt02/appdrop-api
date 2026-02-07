package services

import (
	"appdrop-api/internal/models"
	"appdrop-api/internal/repository"
	"errors"
)

func GetPages() ([]models.Page, error) {
	return repository.GetAllPages()
}

func CreatePage(page models.Page) (*models.Page, error) {

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
