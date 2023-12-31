package subcontroller

import (
	submodel "gin/Subject/subModel"
	"gin/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewSubject(c *gin.Context) {
	var input submodel.AddSubjectRequest
	err := c.ShouldBindJSON(&input)
	if err != nil {
		helpers.ApiFailure(c, http.StatusBadRequest, 1001, err.Error())
		return
	}
	err = input.NewSubValidate()
	if err != nil {
		helpers.ApiFailure(c, http.StatusBadRequest, 2003, err.Error())
		return
	}
	sub, err := submodel.NewSubject(input)
	if err != nil {
		helpers.ApiFailure(c, http.StatusBadRequest, 2003, err.Error())
		return
	}
	helpers.ApiSuccess(c, http.StatusOK, 1002, sub)
}

func UpdateSubject(c *gin.Context) {
	var input submodel.UpdateSubjectRequest
	err := c.ShouldBindJSON(&input)
	if err != nil {
		helpers.ApiFailure(c, http.StatusBadRequest, 2001, err.Error())
		return
	}
	sub, err := submodel.UpdateSubject(input.ID, input)
	if err != nil {
		helpers.ApiFailure(c, http.StatusBadRequest, 2001, err.Error())
		return
	}
	helpers.ApiSuccess(c, http.StatusOK, 1003, sub)
}
