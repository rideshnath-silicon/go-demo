package main


func main() {
	// insert()
	dynamicInsert()
	// getSingle()
	// multi()
	GetAllUsers()
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
