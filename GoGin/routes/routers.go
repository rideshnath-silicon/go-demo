package routes

import (
	"gin/controller"
	"gin/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	r.GET("/", middleware.Login)
	r.GET("/user", middleware.AuthMiddleware(), controller.GetUser)
	r.GET("/users", middleware.AuthMiddleware(), controller.GetAllUser)
	r.POST("/New-user", middleware.AuthMiddleware(), controller.NewUser)
}
