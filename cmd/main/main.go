package main

import (
	"fmt"
	"github.com/rs/cors"
	"log"
	"net/http"
	"test-project-backend/pkg/config"
	"test-project-backend/pkg/controllers"
	"test-project-backend/pkg/database"

	"github.com/gorilla/mux"
)

func main() {
	// Load Configurations
	config.LoadAppConfiguration()

	// Initialize Database
	database.Connect(config.AppConfig.ConnectionString)
	database.Migrate()

	// Initialize the router
	router := mux.NewRouter().StrictSlash(true)

	// Register Routes
	controllers.InitializeRoutes(router)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: true,
	})
	handler := c.Handler(router)

	// Start the server
	log.Println(fmt.Sprintf("Starting Server on %s:%s/v1/customers", "http://localhost", config.AppConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", config.AppConfig.Port), handler))
}
