package models

import (
	"hmvcstructure/config"
)

func GetEmployee(id uint) Employee {
	var employee Employee
	result := config.DB.Where("ID = ?", id).Find(&employee)
	if result.Error != nil {
		panic("failed to retrieve user: " + result.Error.Error())
	}
	return employee
}

func GetAllEmployee() []Employee {
	var employees []Employee

	result := config.DB.Order("id desc").Find(&employees)
	if result.Error != nil {
		panic("Faild to retrieve data : " + result.Error.Error())
	}
	return employees
}

func CreateEmp(data Employee) interface{} {
	var emp = data
	// fmt.Println(emp)
	result := config.DB.Create(&emp)
	if result.Error != nil {
		panic("Faild to retrieve data : " + result.Error.Error())
	}
	return emp

}

func UpdateEmp(id uint, data Employee) Employee {
	var upEmp Employee
	result := config.DB.Model(&upEmp).Where("ID = ?", id).Updates(data)
	if result.Error != nil {
		panic("Faild to retrieve data : " + result.Error.Error())
	}
	return upEmp
}

func RemoveEmp(id uint) Employee {
	var DlEmp Employee
	result := config.DB.Where("ID = ?", id).Delete(&DlEmp)
	if result.Error != nil {
		panic("failed to retrieve user: " + result.Error.Error())
	}
	return DlEmp
}
