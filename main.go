package main

import (
	"github.com/joho/godotenv"
	"github.com/khildamaililhaq/gorest-boilerplate/config"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//load db connection from config/database
	config.ConnectDb(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"))

	// Load Rooutes from config/routes
	config.Routes()
}
