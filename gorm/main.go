package main

var db = con()

func main() {

	err := db.AutoMigrate(&Employee{}, &User{})
	if err != nil {
		panic("failed to perform migrations: " + err.Error())
	}

	// InsertData()
	// NewUser()
	// singelRow()
	multipleData()
}
