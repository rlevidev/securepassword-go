package main

import (
	"log"
	"net/http"

	"securepassword-go/internal/handlers"
)

func main() {
	http.HandleFunc("/validate-password", handlers.ValidatePasswordHandler)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
