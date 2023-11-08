package models

import (
	"gin/config"
)

func LoginUser(email string, pass string) User {
	var user User
	result := config.DB.Where("email = ? ", email).Where("password = ?", pass).Find(&user)
	if result.Error != nil {
		panic("Error:-" + result.Error.Error())
	}
	return user
}

func GetAllUser() []User {
	var user []User
	result := config.DB.Find(&user)
	if result.Error != nil {
		panic("Error:-" + result.Error.Error())
	}
	return user
}

func GetUser(id uint) User {
	var user User
	result := config.DB.Where("id = ?",id).Find(&user)
	if result.Error != nil {
		panic("Error:-" + result.Error.Error())
	}
	return user
}
func NewUser(data User) User {
	var user = data
	result := config.DB.Create(&user)
	if result.Error != nil {
		panic("Error:-" + result.Error.Error())
	}
	return user
}
