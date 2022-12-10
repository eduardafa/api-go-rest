package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDatabase() {
	connectionString := "host=localhost user=root password=root dbname=root port=49153 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connectionString))
	if err != nil {
		log.Panic("Error to connect with database")
	}
}
