package routes

import (
	"go-user-crud/controllers"

	"github.com/gorilla/mux"
)

func ResgiterUserRoute(router *mux.Router) {
	router.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/allusers", controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.GetOneUser).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.UpdateOneUser).Methods("PUT")
	router.HandleFunc("/users/{id}", controllers.DeleteUserById).Methods("DELETE")
}
