package main

import (
	. "fmt"
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvs() {

	err := godotenv.Load()
	if err != nil {
		Println(err)
		log.Fatalf("Error loading .env file")
	}

}
