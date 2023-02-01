package main

import (
	"log"

	"github.com/Kkmikaze/go-rest-api/db"
	"github.com/Kkmikaze/go-rest-api/domain/user/model"
)

func main() {
	err := db.DB.AutoMigrate(&model.User{})

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Migration has been done")
}
