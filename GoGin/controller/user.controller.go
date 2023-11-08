package controller

import (
	"gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUser(r *gin.Context) {
	data := models.GetAllUser()
	r.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "Data": data})
}

func GetUser(r *gin.Context) {
	var input models.User
	err := r.ShouldBindJSON(&input)
	if err != nil {
		r.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data := models.GetUser(input.ID)
	r.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "Data": data})

}

func NewUser(r *gin.Context) {

	var input models.User
	err := r.ShouldBindJSON(&input)
	if err != nil {
		r.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data := models.NewUser(input)
	r.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "Data": data})

}
