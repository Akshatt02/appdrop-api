package main

import (
	"fmt"
	"net/http"
	"os"

	"appdrop-api/internal/db"
	"appdrop-api/internal/handlers"

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

	port := os.Getenv("PORT")
	fmt.Println("Server running on :" + port)
	http.ListenAndServe(":"+port, nil)
}
