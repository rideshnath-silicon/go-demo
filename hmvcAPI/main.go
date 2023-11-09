package main

import (
	"hmvcstructure/config"
	emp "hmvcstructure/employee/models"
	"hmvcstructure/routes"
	user "hmvcstructure/user/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func init() {
	config.ConnectDatabase()
}

func main() {

	err := config.DB.AutoMigrate(&emp.Employee{}, &user.User{})
	if err != nil {
		panic("failed to perform migrations: " + err.Error())
	}
	r := mux.NewRouter()
	routes.IndexRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":5555", nil))
}
