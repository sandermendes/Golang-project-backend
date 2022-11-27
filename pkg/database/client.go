package database

import (
	"gorm.io/driver/postgres"
	"log"
	"test-project-backend/pkg/entities"

	"gorm.io/gorm"
)

var Instance *gorm.DB
var err error

func Connect(connectionString string) {
	// Initiate the database connection
	Instance, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	// Check for error in connection
	if err != nil {
		panic("Cannot connect to DB: " + err.Error())
	}
	log.Println("Connected to Database...")
}

func Migrate() {
	// Create entity in database
	Instance.AutoMigrate(&entities.Customer{})

	log.Println("Database Migration Completed...")
}
