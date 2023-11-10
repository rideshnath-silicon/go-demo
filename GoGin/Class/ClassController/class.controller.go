package classcontroller

import (
	classmodel "gin/Class/ClassModel"
	"gin/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewClass(c *gin.Context) {
	var input classmodel.NewClassRequest
	err := c.ShouldBindJSON(&input)
	if err != nil {
		helpers.ApiFailure(c, http.StatusBadRequest, 1001, err.Error())
		return
	}
	err = input.NewClassValidate()
	if err != nil {
		helpers.ApiFailure(c, http.StatusBadRequest, 2003, err.Error())
		return
	}
	data, err := classmodel.NewClass(input)
	if err != nil {
		helpers.ApiFailure(c, http.StatusBadRequest, 2003, err.Error())
		return
	}
	helpers.ApiSuccess(c, http.StatusOK, 1002, data)
}

func UpdateSubject(c *gin.Context) {
	var input classmodel.UpdateClassRequest
	err := c.ShouldBindJSON(&input)
	if err != nil {
		helpers.ApiFailure(c, http.StatusBadRequest, 2001, err.Error())
		return
	}
	data, err := classmodel.UpdateClass(input.ID, input)
	if err != nil {
		helpers.ApiFailure(c, http.StatusBadRequest, 2001, err.Error())
		return
	}
	helpers.ApiSuccess(c, http.StatusOK, 1003, data)
}
