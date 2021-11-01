package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

var database *Database

func GetDbInstance() *Database {
	if database == nil {
		db, err := gorm.Open(postgres.Open(os.Getenv("DB_DSN")), &gorm.Config{})

		database = &Database{db}
		if err != nil {
			log.Fatal("Failed to open a DB connection: ", err)
		}
	}

	return database
}
