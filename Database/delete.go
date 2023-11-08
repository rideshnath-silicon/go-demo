package main

func delete() {
	db := con()
	insertStmt := `DELETE FROM users WHERE id = $1;`
	_, e := db.Exec(insertStmt, 2)
	CheckError(e)

}

func Update() {
	db := con()
	insertStmt := `UPDATE emp SET "age" = 3000 WHERE id = $1;`
	_, e := db.Exec(insertStmt, 2)
	CheckError(e)

}
