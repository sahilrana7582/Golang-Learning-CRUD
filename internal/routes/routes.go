package routes

import (

	"github.com/gorilla/mux"
	"simple-crud/internal/handlers"

)

func SetupRoutes(router *mux.Router, userHandler *handlers.UserHandler){

	router.HandleFunc("/api/users", userHandler.CreateUser).Methods("POST")
	router.HandleFunc("/api/users/{id}", userHandler.GetUserByID).Methods("GET")

}

