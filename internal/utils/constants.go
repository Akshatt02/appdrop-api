// Package utils provides utility functions and constants used throughout the application.
package utils

// ValidWidgetTypes defines the set of allowed widget types that can be used in the application.
// Only widgets with types in this map are accepted by the API.
// Valid types:
//   - banner: Full-width promotional content with images
//   - product_grid: Grid layout for displaying products
//   - text: Plain or formatted text content
//   - image: Individual image display
//   - spacer: Empty space for layout purposes
var ValidWidgetTypes = map[string]bool{
	"banner":       true,
	"product_grid": true,
	"text":         true,
	"image":        true,
	"spacer":       true,
}
