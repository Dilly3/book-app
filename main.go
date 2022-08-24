package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/dilly3/book-app/routes"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := routes.MountServer()

	r.Run(":" + port)
}
