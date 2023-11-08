package models

import "hmvcstructure/config"

func LoginUser(email string, pass string) User {
	var user User
	result := config.DB.Where("email = ?", email).Where("password = ?", pass).Find(&user)
	if result.Error != nil {
		panic("Error to fetch data" + result.Error.Error())
	}
	return user
}

func Getuser(id uint) User {
	var user User
	result := config.DB.Where("id = ?", id).Find(&user)

	if result.Error != nil {
		panic("Error to fetch data" + result.Error.Error())
	}
	return user
}
func GetAllUser() []User {
	var users []User
	result := config.DB.Order("id desc").Find(&users)
	if result.Error != nil {
		panic("Faild to retrieve data : " + result.Error.Error())
	}
	return users
}

func CreateUser(data User) interface{} {
	var user = data
	// fmt.Println(emp)
	result := config.DB.Create(&user)
	if result.Error != nil {
		panic("Faild to retrieve data : " + result.Error.Error())
	}
	return user

}

func UpdateUser(id uint, data User) User {
	var user User
	result := config.DB.Model(&user).Where("ID = ?", id).Updates(data)
	if result.Error != nil {
		panic("Faild to retrieve data : " + result.Error.Error())
	}
	return user
}
