package routes

import (
	"net/http"
	"restApi/controllers"
)

func HandleRequests() {
	http.HandleFunc("/", controllers.GetEmployee)
	http.HandleFunc("/AllEmployee", controllers.GetAllEmployee)
	http.HandleFunc("/UpdateEmployee", controllers.UpdateEmp)
	http.HandleFunc("/CreateEmployee", controllers.CreateEmp)
	http.HandleFunc("/DeleteEmployee", controllers.DeleteEmp)
}
