package routes

import (
	classroute "gin/Class/ClassRoute"
	studentroute "gin/Student/StudentRoute"
	subroute "gin/Subject/subRoute"
	"gin/User/routes"
	"gin/middleware"

	"github.com/gin-gonic/gin"
)

// func UserRoutes(r *gin.Engine) {
// 	r.GET("/", middleware.Login)
// 	r.GET("/user", middleware.AuthMiddleware(), controller.GetUser)
// 	r.GET("/users", middleware.AuthMiddleware(), controller.GetAllUser)
// 	r.POST("/New-user", middleware.AuthMiddleware(), controller.NewUser)
// 	r.PUT("/UP-user", middleware.AuthMiddleware(), controller.UpdateUser)
// }

func IndexRouter(r *gin.Engine) {
	r.GET("/", middleware.Login)
	r.Use(middleware.AuthMiddleware())
	routes.UserRoutes(r)
	studentroute.StudentRoutes(r)
	subroute.SubRoute(r)
	classroute.Classroute(r)
}
