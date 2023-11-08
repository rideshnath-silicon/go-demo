package routes

import (
	"hmvcstructure/employee/controllers"

	"github.com/gorilla/mux"
)

func EmpRoutes(r *mux.Router) {
	r.HandleFunc("/Emp", controllers.GetEmployee).Methods("GET")
	r.HandleFunc("/Employees", controllers.GetAllEmployee).Methods("GET")
	r.HandleFunc("/UpdateEmployee", controllers.UpdateEmp).Methods("PUT")
	r.HandleFunc("/CreateEmployee", controllers.CreateEmp).Methods("POST")
	r.HandleFunc("/DeleteEmployee", controllers.DeleteEmp).Methods("DELETE")

}
