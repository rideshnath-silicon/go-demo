package routes

import (
	"hmvcstructure/employee/controllers"
	"hmvcstructure/middleware"

	"github.com/gorilla/mux"
)

func EmpRoutes(r *mux.Router) {
	r.HandleFunc("/Emp", middleware.AuthMiddleware(controllers.GetEmployee)).Methods("GET")
	r.HandleFunc("/Employees", middleware.AuthMiddleware(controllers.GetAllEmployee)).Methods("GET")
	r.HandleFunc("/UpdateEmployee", middleware.AuthMiddleware(controllers.UpdateEmp)).Methods("PUT")
	r.HandleFunc("/CreateEmployee", middleware.AuthMiddleware(controllers.CreateEmp)).Methods("POST")
	r.HandleFunc("/DeleteEmployee", middleware.AuthMiddleware(controllers.DeleteEmp)).Methods("DELETE")

}
