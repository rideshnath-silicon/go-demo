package main

import "fmt"

func multi(){
	db := con()

	getALl := `Select * from accounts`

	AllData, err := db.Query(getALl)

	for AllData.Next(){
		var user_id, username, password,email,created_on string
		err = AllData.Scan(&user_id,&username,&password,&email,&created_on)
		CheckError(err)
		fmt.Println(user_id,username, password,email,created_on)
	}
}

func GetAllUsers() {
	db := con()
	getALl := `Select * from users`

	AllData, err := db.Query(getALl)

	for AllData.Next() {
		var id, age, username, last_name, email string
		err = AllData.Scan(&id,&age, &username, &last_name, &email)
		CheckError(err)
		fmt.Println(id,age, username, last_name, email)
	}
}

// INSERT INTO users (age, email, first_name, last_name)
// VALUES (15, 'jerryjr123@gmail.com', 'Jerry', 'Seinfeld');
