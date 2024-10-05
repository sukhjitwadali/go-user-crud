package routes

import (
	"go-user-crud/controllers"

	"github.com/gorilla/mux"
)

func ResgiterUserRoute(router *mux.Router) {
	router.HandleFunc("/users", controllers.CreateUser).Methods("POST")
}
