package main

import (
	"log"

	"github.com/Kkmikaze/go-rest-api/config/db"
	ea "github.com/Kkmikaze/go-rest-api/internal/domain/article/entity"
	ec "github.com/Kkmikaze/go-rest-api/internal/domain/comment/entity"
	eu "github.com/Kkmikaze/go-rest-api/internal/domain/user/entity"
)

func main() {
	err := db.DB.AutoMigrate(&eu.User{})

	if err != nil {
		log.Fatal(err)
	}

	err = db.DB.AutoMigrate(&ea.Article{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.DB.AutoMigrate(&ec.Comment{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Migration has been done")
}
