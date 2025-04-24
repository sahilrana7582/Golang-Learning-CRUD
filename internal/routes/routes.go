package routes

import (

	"github.com/gorilla/mux"
	"simple-crud/internal/handlers"

)

func SetupRoutes(router *mux.Router, userHandler *handlers.UserHandler){

	router.HandleFunc("/api/users", userHandler.CreateUser).Methods("POST")
	router.HandleFunc("/api/users/{id}", userHandler.GetUserByID).Methods("GET")
	router.HandleFunc("/api/users/{id}", userHandler.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/users/{id}", userHandler.DeleteUser).Methods("DELETE")
	router.HandleFunc("/api/users", userHandler.GetAllUsers).Methods("GET")

}

