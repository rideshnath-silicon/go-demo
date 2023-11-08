package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "mydb"
)

func con() *gorm.DB {

	//Create a new Postgresql database connection
	postgresqlDbInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	fmt.Println(postgresqlDbInfo)
	// Open a connection to the database
	db, err := gorm.Open(postgres.Open(postgresqlDbInfo), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}

	return db

}
