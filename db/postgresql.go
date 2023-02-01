package db

import (
	"fmt"

	"github.com/Kkmikaze/go-rest-api/lib/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Postgresql() (*gorm.DB, error) {
	return gorm.Open(postgres.Open(dsn()), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
	})
}

func dsn() string {
	host := "host=" + env.String("DB_HOST", "127.0.0.1")
	port := "port=" + env.String("DB_PORT", "5432")
	dbname := "dbname=" + env.String("DB_NAME", "go_rest_api")
	user := "user=" + env.String("DB_USER", "postgres")
	password := "password=" + env.String("DB_PASSWORD", "postgres")
	return fmt.Sprintln(host, port, dbname, user, password)
}
