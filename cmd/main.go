package main

import (
	"invoicegen/internal/server"
	"log"
	"os"
)

func main() {
	host := os.Getenv("HOST")
	if host == "" {
		host = ":3000"
	}

	dev := os.Getenv("ENV") == "dev"

	dbUri := os.Getenv("DB_URI")
	if dbUri == "" {
		dbUri = "invoicegen.db"
	}

	err := server.Run(&server.Config{
		Host:     host,
		Dev:      dev,
		Database: dbUri,
	})

	if err != nil {
		log.Fatal(err)
	}
}
