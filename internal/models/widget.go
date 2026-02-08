package models

import "time"

// Widget represents a UI component or content block placed on a page.
// Each widget has a specific type and flexible JSON configuration
// that defines its appearance and behavior.
type Widget struct {
	// ID is a UUID that uniquely identifies the widget
	ID string `json:"id"`
	// PageID is the UUID of the page this widget belongs to
	PageID string `json:"page_id"`
	// Type specifies the widget category (banner, product_grid, text, image, spacer)
	Type string `json:"type"`
	// Position is the order index of this widget on its page (0-based)
	Position int `json:"position"`
	// Config holds widget-specific configuration as JSON
	// Structure varies by widget type, e.g., banner has image_url, text has content
	Config map[string]interface{} `json:"config"`
	// CreatedAt is the timestamp when the widget was created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt is the timestamp when the widget was last modified
	UpdatedAt time.Time `json:"updated_at"`
}
