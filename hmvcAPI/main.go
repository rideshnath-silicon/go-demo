package main

import (
	"hmvcstructure/config"
	emp "hmvcstructure/employee/models"
	empRoute "hmvcstructure/employee/routes"
	"hmvcstructure/middleware"
	user "hmvcstructure/user/models"
	userRoute "hmvcstructure/user/routes"
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
	empRoute.EmpRoutes(r)
	userRoute.UserRoutes(r)
	http.Handle("/", r)
	r.HandleFunc("/", middleware.Login)
	log.Fatal(http.ListenAndServe(":5555", nil))
}
