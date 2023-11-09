package studentcontroller

import (
	studentmodel "gin/Student/StudentModel"
	"gin/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetStudent(c *gin.Context) {
	var input studentmodel.GetStudentRequest
	c.ShouldBindJSON(&input)
	studentData, err := studentmodel.GetStudent(input.RollNumber)
	if err != nil {
		helpers.ApiFailure(c, http.StatusBadRequest, 1001, err.Error())
		return
	}
	if studentData == 0 {
		helpers.ApiFailure(c, http.StatusBadRequest, 1001, "Data Not Found")
		return
	}
	helpers.ApiSuccess(c, http.StatusOK, 1000, studentData)
}

func RegisterStudent(c *gin.Context) {
	var input studentmodel.CreateStudentRequest
	err := c.ShouldBindJSON(&input)
	if err != nil {
		helpers.ApiFailure(c, http.StatusBadRequest, 1001, err.Error())
		return
	}
	err = input.NewStudentValidate()
	if err != nil {
		helpers.ApiFailure(c, http.StatusBadRequest, 1001, err.Error())
		return
	}
	studentData, err := studentmodel.RegisterStudent(input)
	if err != nil {
		helpers.ApiFailure(c, http.StatusBadRequest, 1001, err.Error())
		return
	}
	if studentData == 0 {
		helpers.ApiFailure(c, http.StatusBadRequest, 1001, "Data Not Found")
		return
	}
	helpers.ApiSuccess(c, http.StatusOK, 1002, studentData)
}

func UpdateStudent(c *gin.Context) {
	var input studentmodel.UpdateStudentRequest
	err := c.ShouldBindJSON(&input)
	if err != nil {
		helpers.ApiFailure(c, http.StatusBadRequest, 1001, err.Error())
		return
	}
	studentData, err := studentmodel.UpdateStudent(input.ID, input)
	if err != nil {
		helpers.ApiFailure(c, http.StatusBadRequest, 1001, err.Error())
		return
	}
	if studentData == 0 {
		helpers.ApiFailure(c, http.StatusBadRequest, 1001, "Data Not Found")
		return
	}
	helpers.ApiSuccess(c, http.StatusOK, 1003, studentData)
}
