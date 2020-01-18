package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

//AppName : application name
var AppName string

//HTTPPort : rest api port
var HTTPPort string

//GinMode : gin spesific mode
var GinMode string

//DbName : database name
var DbName string

//DbDebug : if set to true will print out the query string
var DbDebug bool

//Load : get .env file then set the environment variables
func Load() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file ...")
		log.Println("Using default configuration")
	}

	if AppName = os.Getenv("APP_NAME"); AppName == "" {
		AppName = "My App"
	}

	if GinMode = os.Getenv("GIN_MODE"); GinMode != "release" {
		GinMode = "debug"
	}

	if HTTPPort = os.Getenv("APP_PORT"); HTTPPort == "" {
		HTTPPort = "8080"
	}

	if DbName = os.Getenv("DB_NAME"); DbName == "" {
		DbName = "mydata"
	}

	if config := os.Getenv("DB_DEBUG"); config == "true" {
		DbDebug = true
	}
}
