package services

import (
	"appdrop-api/internal/models"
	"appdrop-api/internal/repository"
	"errors"
)

func GetPages() ([]models.Page, error) {
	return repository.GetAllPages()
}

func CreatePage(page models.Page) error {

	if page.Name == "" || page.Route == "" {
		return errors.New("name and route are required")
	}

	return repository.CreatePage(page)
}
