package main

import (
	"fmt"
	"log"
	"net/http"
	"test-project-backend/pkg/config"
	"test-project-backend/pkg/controllers"
	"test-project-backend/pkg/database"

	"github.com/gorilla/mux"
)

func main() {

	// Load Configurations from config.json using Viper
	config.LoadAppConfiguration()

	// Initialize Database
	database.Connect(config.AppConfig.ConnectionString)
	database.Migrate()

	// Initialize the router
	router := mux.NewRouter().StrictSlash(true)

	// Register Routes
	controllers.InitializeRoutes(router)

	// Start the server
	log.Println(fmt.Sprintf("Starting Server on port %s:%s", "http://localhost", config.AppConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", config.AppConfig.Port), router))
}