package subroute

import (
	subcontroller "gin/Subject/subController"

	"github.com/gin-gonic/gin"
)

func SubRoute(r *gin.Engine) {
	r.POST("/sub/New", subcontroller.NewSubject)
	r.POST("/sub/Update", subcontroller.UpdateSubject)
}
