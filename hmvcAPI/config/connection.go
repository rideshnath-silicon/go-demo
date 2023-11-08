package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	DbInfo := "host=localhost port=5432 user=postgres password=root dbname=mydb sslmode=disable"

	db, err := gorm.Open(postgres.Open(DbInfo))

	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}

	DB = db
}
