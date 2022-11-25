package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ConnectionString string `mapstructure:"connection_string"`
	Port             string `mapstructure:"port"`
}

var AppConfig *Config

func LoadAppConfiguration() {
	log.Println("Loading Server Configurations...")

	// Set path for configuration file
	viper.AddConfigPath(".")

	// Name file
	viper.SetConfigName("config")

	// File type
	viper.SetConfigType("json")

	// Check if config file can be loaded
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatal(err)
	}
}
