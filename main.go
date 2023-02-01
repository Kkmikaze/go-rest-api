package main

import (
	"log"

	_ "github.com/joho/godotenv/autoload"

	"github.com/Kkmikaze/go-rest-api/config"
)

func main() {
	err := config.Routers.Run()
	if err != nil {
		log.Fatal(err)
	}
}
