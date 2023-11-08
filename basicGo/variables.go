package main

import "fmt"

var str string
var vint int
var unsignint uint = 255 // unsign data type always use positve integer
var float float32
var boolean bool
var x = "2"

const VCONSTENT = 5000

func variables() {

	str = "Ridesh Nath"
	str = "Ridesh Nath Ji"

	var (
		_str1 = "demo"
		vint1 = 500
	)

	fmt.Println(str, " after groupd new str demo ", _str1)
	fmt.Printf("This variable have value = \"%v\" \n and the type of varible is = \"%T\" \n", str, str)
	fmt.Printf("This variable have value = %q \n", str)
	fmt.Printf("This variable have value = %s \n and the type of varible is = \"%T\" \n", str, str)
	fmt.Print(unsignint)
	fmt.Println("\nint value", vint, "\n after grouped new int demo ", vint1)
	fmt.Println(float)
	fmt.Printf("I am constant variable %t  ", boolean)
	fmt.Println(VCONSTENT)
	fmt.Print(x,"\n") // if we not use \n here then the next output not in generate in new line
	fmt.Println("Hello World!")

}
