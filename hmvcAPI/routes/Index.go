package routes

import (
	empRoute "hmvcstructure/employee/routes"
	"hmvcstructure/middleware"
	userRoute "hmvcstructure/user/routes"

	"github.com/gorilla/mux"
)

func IndexRoutes(r *mux.Router) {
	r.HandleFunc("/", middleware.Login)
	empRoute.EmpRoutes(r)
	userRoute.UserRoutes(r)
}
