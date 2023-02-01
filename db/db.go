package db

import (
	"log"

	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	var err error
	DB, err = Postgresql()
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	dbSQL, ok := DB.DB()
	if ok != nil {
		defer dbSQL.Close()
	}
}
