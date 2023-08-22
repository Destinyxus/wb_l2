package main

import (
	"log"

	"github.com/joho/godotenv"

	"dev11/internal/apiserver"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	server := apiserver.NewApiServer()
	if err := server.Run(); err != nil {
		panic(err)
	}

}
