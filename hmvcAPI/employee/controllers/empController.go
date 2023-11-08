package controllers

import (
	"fmt"
	"hmvcstructure/employee/models"
	"hmvcstructure/helper"
	"net/http"
)

func GetEmployee(w http.ResponseWriter, r *http.Request) {
	var emp models.Employee
	bodyData := helper.RequestBody(w, r, emp)
	empID := bodyData.ID
	myData := models.GetEmployee(empID)
	if myData.ID == 0 && myData.Email == "" {
		helper.ApiFailure(w, 200)
		return
	}
	helper.ApiSuccess(w, myData, 100)
}

func GetAllEmployee(w http.ResponseWriter, r *http.Request) {
	myData := models.GetAllEmployee()
	helper.ApiSuccess(w, myData, 100)

}

func CreateEmp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method. Use POST.", http.StatusMethodNotAllowed)
		return
	}
	var emp models.Employee
	bodyData := helper.RequestBody(w, r, emp)
	valERR := validateEmployee(bodyData)
	fmt.Println("Validation failed:", valERR)
	if valERR != nil {
		http.Error(w, valERR.Error(), http.StatusBadRequest)
		return
	}
	crData := models.CreateEmp(bodyData)
	helper.ApiSuccess(w, crData, 102)

}

func UpdateEmp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Invalid request method. Use PUT.", http.StatusMethodNotAllowed)
		return
	}
	var emp models.Employee
	bodyData := helper.RequestBody(w, r, emp)
	Id := bodyData.ID
	myData := models.GetEmployee(Id)
	if myData.ID == 0 && myData.Email == "" {
		helper.ApiFailure(w, 200)
		return
	}
	updateData := models.UpdateEmp(Id, bodyData)
	helper.ApiSuccess(w, updateData, 103)

}

func DeleteEmp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method. Use DELETE.", http.StatusMethodNotAllowed)
		return
	}
	var emp models.Employee
	bodyData := helper.RequestBody(w, r, emp)
	Id := bodyData.ID
	myData := models.GetEmployee(Id)
	if myData.ID == 0 && myData.Email == "" {
		helper.ApiFailure(w, 200)
		return
	}
	// fmt.Println(myData)
	Output := models.RemoveEmp(Id)

	helper.ApiSuccess(w, Output, 102)

}

func validateEmployee(data models.Employee) error {
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
	if data.Age < 0 && data.Age > 999 {
		return fmt.Errorf("age must be non-negative")
	}
	return nil
}
