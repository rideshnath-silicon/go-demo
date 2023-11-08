package controller

import (
	"fmt"
	"hmvcstructure/helper"
	"hmvcstructure/user/models"
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	var user models.User
	bodyData := helper.RequestBody(w, r, user)
	userId := bodyData.ID
	resultData := models.Getuser(userId)
	if resultData.ID == 0 && resultData.Email == "" {
		helper.ApiFailure(w, 200)
		return
	}
	helper.ApiSuccess(w, resultData, 100)

}

func GetAllUser(w http.ResponseWriter, r *http.Request) {

	// w.Write([]byte("Access granted to protected resource"))
	myData := models.GetAllUser()
	helper.ApiSuccess(w, myData, 100)
}

func CreateNewUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method. Use POST.", http.StatusMethodNotAllowed)
		return
	}
	var user models.User
	bodyData := helper.RequestBody(w, r, user)
	valERR := validateUser(bodyData)
	fmt.Println("Validation failed:", valERR)
	if valERR != nil {
		http.Error(w, valERR.Error(), http.StatusBadRequest)
		return
	}
	crData := models.CreateUser(bodyData)
	helper.ApiSuccess(w, crData, 102)

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	if r.Method != http.MethodPut {
		http.Error(w, "Invalid request method. Use PUT.", http.StatusMethodNotAllowed)
		return
	}
	var emp models.User
	bodyData := helper.RequestBody(w, r, emp)
	Id := bodyData.ID
	myData := models.Getuser(Id)
	if myData.ID == 0 && myData.Email == "" {
		helper.ApiFailure(w, 200)
		return
	}
	data := models.UpdateUser(Id, bodyData)
	helper.ApiSuccess(w, data.Email, 103)

}

func validateUser(data models.User) error {
	if data.FirstName == "" {
		return fmt.Errorf("plaese Enter name")
	}
	if data.FirstName == "" {
		return fmt.Errorf("plaese Enter Firstname")
	}
	if data.LastName == "" {
		return fmt.Errorf("plaese Enter Lastname")
	}
	if data.Country == "" {
		return fmt.Errorf("plaese Enter Country")
	}
	if data.Role == "" {
		return fmt.Errorf("plaese Enter Role")
	}
	if data.Age < 0 {
		return fmt.Errorf("age must be non-negative")
	}
	return nil
}
