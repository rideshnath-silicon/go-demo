package routes

import (
	"cardemo-crud/User/controller"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {

	r.GET("/user", controller.GetLoginUser)
	r.GET("/Single-user", controller.GetUser)
	r.GET("/users", controller.GetAllUser)
	r.POST("/New-user", controller.NewUser)
	r.PUT("/UP-user", controller.UpdateUser)
	r.DELETE("/Delete-User", controller.DeleteUser)
}
