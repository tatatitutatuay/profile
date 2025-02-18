package main

import (
	"log"
	"net/http"
	
	"profile-backend/config"
	"profile-backend/internal/controllers"
	"github.com/joho/godotenv"
	"github.com/gorilla/mux"
)

func main() {
	// Initialize MongoDB
	config.ConnectMongoDB()

	// Define routes
	router := mux.NewRouter()
	router.HandleFunc("/profile", controllers.UpdateProfile).Methods("PUT")
	//router.HandleFunc("/projects", controllers.HandleProjects)
	//router.HandleFunc("/experiences", controllers.HandleExperiences)

	// Start server
	//utils.Logger().Println("Server is running on port 8080")
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get port from environment variables
	port := ":" + config.GetEnv("PORT")

	log.Fatal(http.ListenAndServe(port, router))
}