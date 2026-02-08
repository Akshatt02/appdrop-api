// Package main provides the entry point for the AppDrop API server.
// It sets up HTTP routing, middleware, and database connections.
package main

import (
	"fmt"
	"net/http"
	"os"

	"appdrop-api/internal/db"
	"appdrop-api/internal/handlers"
	"appdrop-api/internal/middleware"

	"github.com/joho/godotenv"
)

// main initializes and starts the AppDrop API server.
// It performs the following:
// 1. Loads environment variables from .env file
// 2. Establishes PostgreSQL database connection
// 3. Registers HTTP route handlers for all endpoints
// 4. Applies request logging middleware
// 5. Starts the HTTP server on the configured port
func main() {
	// Load environment variables from .env file for configuration
	godotenv.Load()

	// Initialize database connection pool using configured DATABASE_URL
	db.ConnectDB()

	// Health check endpoint - validates API and database connectivity
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "API + DB working")
	})

	// Pages list and creation endpoints
	// GET /pages - List all pages
	// POST /pages - Create a new page
	http.HandleFunc("/pages", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handlers.GetPagesHandler(w, r)
			return
		}
		if r.Method == http.MethodPost {
			handlers.CreatePageHandler(w, r)
			return
		}
		http.NotFound(w, r)
	})

	// Pages detail, update, delete, and widget operations
	// GET /pages/:id - Get page with all its widgets
	// PUT /pages/:id - Update page details
	// DELETE /pages/:id - Delete page and all its widgets
	// POST /pages/:id/widgets - Create new widget on page
	// POST /pages/:id/widgets/reorder - Reorder widgets on page

	http.HandleFunc("/pages/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		// Handle POST /pages/:id/widgets/reorder
		// Reorders all widgets on a specific page by updating their positions
		if r.Method == http.MethodPost && len(path) > len("/pages/") {
			if len(path) >= len("/widgets/reorder") && path[len(path)-len("/widgets/reorder"):] == "/widgets/reorder" {
				handlers.ReorderWidgetsHandler(w, r)
				return
			}
			// Handle POST /pages/:id/widgets
			// Creates a new widget on the specified page
			if len(path) >= len("/widgets") && path[len(path)-len("/widgets"):] == "/widgets" {
				handlers.CreateWidgetHandler(w, r)
				return
			}
		}

		// Handle GET /pages/:id
		// Retrieves page details including all associated widgets
		if r.Method == http.MethodGet {
			handlers.GetPageByIDHandler(w, r)
			return
		}

		// Handle PUT /pages/:id
		// Updates page name, route, or home page status
		if r.Method == http.MethodPut {
			handlers.UpdatePageHandler(w, r)
			return
		}

		// Handle DELETE /pages/:id
		// Deletes page and cascades delete to all its widgets
		if r.Method == http.MethodDelete {
			handlers.DeletePageHandler(w, r)
			return
		}

		http.NotFound(w, r)
	})

	// Widget update and delete endpoints
	// PUT /widgets/:id - Update widget configuration or position
	// DELETE /widgets/:id - Delete a widget from its page

	http.HandleFunc("/widgets/", func(w http.ResponseWriter, r *http.Request) {
		// Handle PUT /widgets/:id
		// Updates widget configuration, position, or type
		if r.Method == http.MethodPut {
			handlers.UpdateWidgetHandler(w, r)
			return
		}
		// Handle DELETE /widgets/:id
		// Removes a widget from its page
		if r.Method == http.MethodDelete {
			handlers.DeleteWidgetHandler(w, r)
			return
		}
		http.NotFound(w, r)
	})

	// Start HTTP server with logging middleware
	// Read port from environment variable (default: 8080)
	// Apply request logging middleware to track all API calls
	port := os.Getenv("PORT")
	fmt.Println("Server running on :" + port)
	handler := middleware.Logger(http.DefaultServeMux)
	http.ListenAndServe(":"+port, handler)
}
