package main


import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"simple-crud/pkg/database"
	"log"
	"simple-crud/internal/handlers"
	"simple-crud/internal/repository"
	"simple-crud/internal/routes"

)

func main() {

	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}

	db, err := database.InitDB()
	if err != nil {
		fmt.Println("Error initializing database:", err)
		return
	}
	defer db.Close()
	if err := database.CreateTables(db); err != nil {
        fmt.Println("Error creating tables:", err)
        return
    }

	userRepository := repository.NewUserRepository(db)
	userHandler := handlers.NewUserHandler(userRepository)


	router := mux.NewRouter()



	routes.SetupRoutes(router, userHandler)
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "OK")
	}).Methods("GET")

	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
        log.Fatal(err)
    }

}