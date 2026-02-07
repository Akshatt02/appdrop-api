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

func main() {
	godotenv.Load()

	db.ConnectDB()

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "API + DB working 🚀")
	})

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

	http.HandleFunc("/pages/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		// POST /pages/:id/widgets/reorder
		if r.Method == http.MethodPost && len(path) > len("/pages/") {
			if len(path) >= len("/widgets/reorder") && path[len(path)-len("/widgets/reorder"):] == "/widgets/reorder" {
				handlers.ReorderWidgetsHandler(w, r)
				return
			}
			// POST /pages/:id/widgets
			if len(path) >= len("/widgets") && path[len(path)-len("/widgets"):] == "/widgets" {
				handlers.CreateWidgetHandler(w, r)
				return
			}
		}

		// GET /pages/:id
		if r.Method == http.MethodGet {
			handlers.GetPageByIDHandler(w, r)
			return
		}

		// PUT /pages/:id
		if r.Method == http.MethodPut {
			handlers.UpdatePageHandler(w, r)
			return
		}

		// DELETE /pages/:id
		if r.Method == http.MethodDelete {
			handlers.DeletePageHandler(w, r)
			return
		}

		http.NotFound(w, r)
	})

	http.HandleFunc("/widgets/", func(w http.ResponseWriter, r *http.Request) {
		// PUT /widgets/:id
		if r.Method == http.MethodPut {
			handlers.UpdateWidgetHandler(w, r)
			return
		}
		// DELETE /widgets/:id
		if r.Method == http.MethodDelete {
			handlers.DeleteWidgetHandler(w, r)
			return
		}
		http.NotFound(w, r)
	})

	port := os.Getenv("PORT")
	fmt.Println("Server running on :" + port)
	handler := middleware.Logger(http.DefaultServeMux)
	http.ListenAndServe(":"+port, handler)
}
