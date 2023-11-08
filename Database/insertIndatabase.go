package main

func insert() {
	db := con()
	insertStmt := `insert into "accounts"("username", "password","email","created_on") values('ridesh', 12345,'Rides@gmail.com','2023-11-03')`
	_, e := db.Exec(insertStmt)
	CheckError(e)

}

// INSERT INTO users (age, email, first_name, last_name)
// VALUES (30, 'jon@calhoun.io', 'Jonathan', 'Calhoun');
// INSERT INTO users (age, email, first_name, last_name)
// VALUES (52, 'bob@smith.io', 'Bob', 'Smith');
// INSERT INTO users (age, email, first_name, last_name)
// VALUES (15, 'jerryjr123@gmail.com', 'Jerry', 'Seinfeld');

func dynamicInsert() {
	db := con()
	insertDynStmt := `insert into "users"("age", "email", "first_name", "last_name") values($1, $2,$3,$4)`
	_, e := db.Exec(insertDynStmt, 22, "ridesh@gmail.com","ridesh","nath")
	CheckError(e)
}
