package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	ConnectionString string
	Port             string
}

var AppConfig Config

func LoadAppConfiguration() Config {
	log.Println("Loading Server Configurations...")

	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file. Err: %s", err)
	}

	var AppPort = os.Getenv("APP_LISTEN_PORT")
	if AppPort == "" {
		AppPort = "8080"
	}

	// Compose the connection string
	var DbHost = "host=" + os.Getenv("DB_HOST")
	var DbPort = "port=" + os.Getenv("DB_PORT")
	var DbUser = "user=" + os.Getenv("DB_USER")
	var DbPassword = "password=" + os.Getenv("DB_PASSWORD")
	var DbName = "dbname=" + os.Getenv("DB_DBNAME")
	var DbSslMode = "sslmode=" + os.Getenv("DB_SSLMODE")

	AppConfig.ConnectionString = DbHost + " " + DbPort + " " + DbUser + " " + DbPassword + " " + DbName + " " + DbSslMode
	AppConfig.Port = AppPort

	return AppConfig
}
