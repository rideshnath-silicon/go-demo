package main

import (
	"fmt"
	"strconv"
)

func typeCasting() {
	fmt.Println("Type Casting start here")

	a := 4.56
	b := "5"
	g := 4
	var c float64
	var d int

	d, err := strconv.Atoi(b)

	c = a + float64(g)
	f := strconv.Itoa(g)
	fmt.Printf("The value of the variable C is = %v \n", c)
	fmt.Printf("The value of the variable d is = %v and type %T \n", d, d)
	fmt.Printf("The value of the variable d is = %q and type %T \n", f, f)
	fmt.Println(err)

}
