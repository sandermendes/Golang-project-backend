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

	// Insert some test rows
	Instance.Create(&entities.Customer{
		FirstName: "Sander",
		LastName:  "Mendes",
		Email:     "sandermendes@gmail.com",
	})
	Instance.Create(&entities.Customer{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "test@test.com",
	})
	Instance.Create(&entities.Customer{
		FirstName: "Jane",
		LastName:  "Doe",
		Email:     "test@test.com",
	})
	log.Println("Database Migration Completed...")
}
