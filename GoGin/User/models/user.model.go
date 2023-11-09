package models

import (
	"gin/config"
	"time"
)

func LoginUser(email string, pass string) (User,error) {
	var user User
	result := config.DB.Select("email").Where("email = ? ", email).Where("password = ?", pass).Find(&user)
	if result.Error != nil {
		return user ,result.Error
	}
	return user,nil
}

func GetAllUser() []User {
	var user []User
	result := config.DB.Find(&user)
	if result.Error != nil {
		panic("Error:-" + result.Error.Error())
	}
	return user
}

func GetUser(id uint) interface{} {
	var user User
	result := config.DB.Where("id = ?", id).Find(&user)
	if result.Error != nil {
		return 0
	}
	if result.RowsAffected == 0 {
		return 0
	}
	return user
}
func NewUser(data NewUserRequest) interface{} {
	user := User{
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Age:       data.Age,
		Email:     data.Email,
		Country:   data.Country,
		Password:  data.Password,
		Role:      data.Role,
		CreatedAt: time.Now(),
	}
	result := config.DB.Create(&user)
	if result.Error != nil {
		return result.Error.Error()
	}
	if result.RowsAffected == 0 {
		return 0
	}
	return user
}

func UpdateUser(id uint, data UpdateUserRequest) interface{} {
	var user = User{
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Email:     data.Email,
		Password:  data.Password,
		Role:      data.Role,
		Country:   data.Country,
		Age:       data.Age,
		UpdatedAt: time.Now(),
	}
	result := config.DB.Where("id = ?", id).Updates(user).Model(&user)
	if result.Error != nil {
		return result.Error.Error()
	}
	if result.RowsAffected == 0 {
		return 0
	}
	return data
}
