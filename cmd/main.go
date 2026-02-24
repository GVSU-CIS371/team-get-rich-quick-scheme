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

	err := server.Run(&server.Config{
		Host: host,
		Dev:  dev,
	})

	if err != nil {
		log.Fatal(err)
	}
}
