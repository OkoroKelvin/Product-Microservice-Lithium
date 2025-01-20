package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	// Load environment variables
	log.Println("Loading environment variables...")
	if os.Getenv("GO_ENV") == "local" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
	}
}

func main() {
	log.Println("Starting server...")

	app, err := Bootstrap()
	if err != nil {
		log.Fatalf("Failed to bootstrap app: %v", err)
	}

	// app.StartHttpServer()
	app.StartGRPCServer()
}
