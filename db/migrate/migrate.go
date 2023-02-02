package main

import (
	"log"

	"github.com/Kkmikaze/go-rest-api/db"
	modelArticle "github.com/Kkmikaze/go-rest-api/domain/article/model"
	modelUser "github.com/Kkmikaze/go-rest-api/domain/user/model"
)

func main() {
	err := db.DB.AutoMigrate(&modelUser.User{})

	if err != nil {
		log.Fatal(err)
	}

	err = db.DB.AutoMigrate(&modelArticle.Article{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Migration has been done")
}
