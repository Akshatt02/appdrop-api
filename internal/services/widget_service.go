package services

import (
	"errors"

	"appdrop-api/internal/models"
	"appdrop-api/internal/repository"
	"appdrop-api/internal/utils"
)

// CreateWidget validates and creates a new widget on a page.
// Business Rules Enforced:
//   - Widget type must be one of the valid types: banner, product_grid, text, image, spacer
//   - Page specified by PageID must exist
//
// Returns the created widget with its UUID or an error.
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

// UpdateWidget modifies an existing widget's properties.
// Business Rules Enforced:
//   - Widget type must be one of the valid types: banner, product_grid, text, image, spacer
//   - Widget must exist by ID
//
// Returns the updated widget or an error.
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

// DeleteWidget removes a widget from the database.
// Validates the widget exists before deletion.
// Returns error if widget not found.
func DeleteWidget(id string) error {
	// Validate widget exists
	_, err := repository.GetWidgetByID(id)
	if err != nil {
		return errors.New("widget not found")
	}

	return repository.DeleteWidget(id)
}

// ReorderWidgets updates the position of all widgets on a page.
// Business Rules Enforced:
//   - Page must exist
//   - All widget IDs must exist
//   - All widgets must belong to the specified page
//
// The position is determined by the order in the ids array (0-based indexing).
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
