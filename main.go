package main

import (
	"fmt"
	"net/http"
	"os"

	"appdrop-api/internal/db"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	db.ConnectDB()

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "API + DB working 🚀")
	})

	port := os.Getenv("PORT")
	fmt.Println("Server running on :" + port)
	http.ListenAndServe(":"+port, nil)
}
