package routes

import (
	"hmvcstructure/middleware"
	"hmvcstructure/user/controller"

	"github.com/gorilla/mux"
)

var Router = mux.NewRouter()

func UserRoutes(r *mux.Router) {

	r.HandleFunc("/user", middleware.AuthMiddleware(controller.GetUser)).Methods("GET")
	r.HandleFunc("/Alluser", middleware.AuthMiddleware(controller.GetAllUser)).Methods("GET")
	r.HandleFunc("/Newuser", middleware.AuthMiddleware(controller.CreateNewUser)).Methods("POST")
	r.HandleFunc("/Upuser", middleware.AuthMiddleware(controller.UpdateUser)).Methods("PUT")
}
