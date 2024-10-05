package main

import (
	"go-user-crud/routes" // Import your route handling package
	"log"
	"net/http"

	"github.com/gorilla/mux" // Import the Gorilla Mux package
)

func main() {
	// Create a new router of type *mux.Router
	router := mux.NewRouter()

	// Register your user routes
	routes.ResgiterUserRoute(router)

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router)) // Start the server
}
