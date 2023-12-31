package studentroute

import (
	studentcontroller "gin/Student/StudentController"

	"github.com/gin-gonic/gin"
)

func StudentRoutes(r *gin.Engine) {
	r.GET("/student", studentcontroller.GetStudent)
	r.GET("/student/class", studentcontroller.GetClassWiseStudent)
	r.POST("/New-student", studentcontroller.RegisterStudent)
	r.PUT("/UP-student", studentcontroller.UpdateStudent)
}
