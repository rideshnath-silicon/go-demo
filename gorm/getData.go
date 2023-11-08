package main

import "fmt"

func singelRow() {

	var user Employee
	result := db.Select("FirstName", "LastName").Where("ID = ?", 4). /*First(&user)*/ Find(&user).Limit(2)
	// result := db.Last(&user)
	if result.Error != nil {
		panic("failed to retrieve user: " + result.Error.Error())
	}
	fmt.Println(user)
	// Use the user record
	// fmt.Printf("User ID: %d, Name: %s %s, Email: %s\n", user.ID,user.FirstName, user.LastName, user.Email)
}

func multipleData() {
	var users []Employee
	result := db.Find(&users)
	if result.Error != nil {
		panic("failed to retrieve user: " + result.Error.Error())
	}

	
	for i := 0; i < len(users); i++ {
		user := users[i]
		fmt.Printf("User ID: %d, Name: %s %s, Email: %s\n", user.ID,
			user.FirstName, user.LastName, user.Email)
	}

}
