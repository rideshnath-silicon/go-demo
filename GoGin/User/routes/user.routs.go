package routes

import (
	"gin/User/controller"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {

	r.GET("/user", controller.GetUser)

	r.GET("/users", controller.GetAllUser)

	r.POST("/New-user", controller.NewUser)

	r.PUT("/UP-user", controller.UpdateUser)
}
