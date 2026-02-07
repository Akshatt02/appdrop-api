package services

import (
	"errors"

	"appdrop-api/internal/models"
	"appdrop-api/internal/repository"
	"appdrop-api/internal/utils"
)

func CreateWidget(widget models.Widget) (*models.Widget, error) {

	if !utils.ValidWidgetTypes[widget.Type] {
		return nil, errors.New("invalid widget type")
	}

	// Validate page exists
	_, err := repository.GetPageByID(widget.PageID)
	if err != nil {
		return nil, errors.New("page not found")
	}

	return repository.CreateWidget(widget)
}

func UpdateWidget(widget models.Widget) (*models.Widget, error) {

	if !utils.ValidWidgetTypes[widget.Type] {
		return nil, errors.New("invalid widget type")
	}

	// Validate widget exists
	_, err := repository.GetWidgetByID(widget.ID)
	if err != nil {
		return nil, errors.New("widget not found")
	}

	return repository.UpdateWidget(widget)
}

func DeleteWidget(id string) error {
	// Validate widget exists
	_, err := repository.GetWidgetByID(id)
	if err != nil {
		return errors.New("widget not found")
	}

	return repository.DeleteWidget(id)
}

func ReorderWidgets(pageID string, ids []string) error {
	// Validate page exists
	_, err := repository.GetPageByID(pageID)
	if err != nil {
		return errors.New("page not found")
	}

	// Validate all widgets exist and belong to page
	for _, id := range ids {
		widget, err := repository.GetWidgetByID(id)
		if err != nil {
			return errors.New("widget not found")
		}
		if widget.PageID != pageID {
			return errors.New("widget does not belong to this page")
		}
	}

	return repository.ReorderWidgets(pageID, ids)
}
