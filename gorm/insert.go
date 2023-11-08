package main

import "fmt"

func InsertData() {
	newUser := Employee{
		FirstName: "Dev",
		LastName:  "pohekar",
		Email:     "devb@gmail.com",
		Country:   "India",
		Role:      "Engineer",
		Age:       24,
	}

	// ... Create a new user record...
	result := db.Create(&newUser)
	if result.Error != nil {
		panic("failed to create user: " + result.Error.Error())
	}
	// ... Handle successful creation ...
	fmt.Printf("New user %s %s was created successfully!\n", newUser.FirstName, newUser.LastName)
}

func NewUser(){
	user :=User{
		FirstName: "Ridesh",
		LastName: "Nath",
		Email: "Rideshnath@gmail.com",
		Country: "India",
		Role: "Developer",
		Age: 23,
	}

	create := db.Create(&user)
	
	if create.Error != nil {
		panic("failed to create user: " + create.Error.Error())

	}
	fmt.Printf("New user %s %s was created successfully!\n", user.FirstName, user.LastName)

}