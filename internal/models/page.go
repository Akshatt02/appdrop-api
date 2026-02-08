// Package models defines the core data structures used throughout the AppDrop API.
package models

import "time"

// Page represents a screen or view in a mobile application.
// Each page has a unique route and can contain multiple widgets.
// Only one page can be designated as the home page (is_home=true).
type Page struct {
	// ID is a UUID that uniquely identifies the page
	ID string `json:"id"`
	// Name is the human-readable title of the page
	Name string `json:"name"`
	// Route is the unique URL path for accessing this page (e.g., "/home", "/products")
	Route string `json:"route"`
	// IsHome indicates if this is the application's home/default page
	IsHome bool `json:"is_home"`
	// CreatedAt is the timestamp when the page was created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt is the timestamp when the page was last modified
	UpdatedAt time.Time `json:"updated_at"`
}
