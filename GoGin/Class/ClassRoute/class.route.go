package classroute

import (
	classcontroller "gin/Class/ClassController"

	"github.com/gin-gonic/gin"
)

func Classroute(r *gin.Engine) {
	r.POST("/class/new", classcontroller.NewClass)
	r.POST("/class/update", classcontroller.UpdateSubject)
}
