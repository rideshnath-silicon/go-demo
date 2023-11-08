package config

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	DbInfo := os.Getenv("DB_CON")
	db, err := gorm.Open(postgres.Open(DbInfo))
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}
	DB = db
}
