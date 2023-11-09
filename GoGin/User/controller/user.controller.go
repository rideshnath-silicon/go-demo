package controller

import (
	"gin/User/models"
	"gin/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUser(r *gin.Context) {
	data := models.GetAllUser()
	helpers.ApiSuccess(r, http.StatusOK, 1000, data)
}

func GetUser(r *gin.Context) {
	var input models.GetUserRequest
	err := r.ShouldBindJSON(&input)
	if err != nil {
		helpers.ApiFailure(r, http.StatusBadRequest, 1001, err.Error())
		return
	}
	valErr := input.UserIdValidate()
	if valErr != nil {
		helpers.ApiFailure(r, http.StatusBadRequest, 1001, valErr.Error())
		return
	}
	data := models.GetUser(input.ID)
	if data == 0 {
		helpers.ApiFailure(r, http.StatusBadRequest, 1001, "Error in Get data")
		return
	}
	helpers.ApiSuccess(r, http.StatusOK, 1000, data)

}

func NewUser(r *gin.Context) {
	var input models.NewUserRequest
	err := r.ShouldBindJSON(&input)
	if err != nil {
		// r.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		helpers.ApiFailure(r, http.StatusBadRequest, 1001, err.Error())
		return
	}
	valErr := input.NewUserValidate()
	if valErr != nil {
		helpers.ApiFailure(r, http.StatusBadRequest, 1001, valErr.Error())
		return
	}
	data := models.NewUser(input)
	if data == 0 {
		helpers.ApiFailure(r, http.StatusBadRequest, 1001, "Error in Insert data")
	}
	helpers.ApiSuccess(r, http.StatusOK, 1000, data)
}

func UpdateUser(r *gin.Context) {
	var input models.UpdateUserRequest
	err := r.ShouldBindJSON(&input)
	if err != nil {
		helpers.ApiFailure(r, http.StatusBadRequest, 1001, err.Error())
		return
	}
	data := models.UpdateUser(input.ID, input)
	if data == 0 {
		helpers.ApiFailure(r, http.StatusBadRequest, 1001, "Error in update data")
	}
	helpers.ApiSuccess(r, http.StatusOK, 1003, data)
}
