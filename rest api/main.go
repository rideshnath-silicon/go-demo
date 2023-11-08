package main

import (
	"log"
	"net/http"
	"restApi/config"
	"restApi/models"
	"restApi/routes"
)

func init() {
	config.ConnectDatabase()
}

func main() {

	err := config.DB.AutoMigrate(&models.Employee{}, &models.User{})
	if err != nil {
		panic("failed to perform migrations: " + err.Error())
	}
	routes.HandleRequests()
	log.Fatal(http.ListenAndServe(":10000", nil))
}
