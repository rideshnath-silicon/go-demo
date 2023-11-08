package main

import "fmt"

func getSingle() {
	db := con()
	get := `SELECT username,email FROM accounts where "user_id" = $1` // $1 is a parameter we can pass in exicu and query and queryRow 
	data := db.QueryRow(get,5)
	var name string
	var email string
	data.Scan(&name, &email)

	fmt.Println(name, email)
}
