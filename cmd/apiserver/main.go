package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/vlasove/test05/internal/app/apiserver"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env files found")
	}
}

func main() {
	config := apiserver.NewConfig()

	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
