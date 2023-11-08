package main

import (
	"errors"
	"fmt"
)

var errs = errors.New("Can't Devide")

func errhandl() {
	var x int = 5
	var y int = 2
	out, err := divide(x, y)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(out)
	}

	forErrorHandling()
}

func divide(a int, b int) (int, error) {
	defer HandlePanic()
	if b == 0 {
		panic(errs)
		//  return 0, devideError; // This code for you dont want any return from this function afer genaration err
	}
	return a / b, nil
}

func HandlePanic() {
	r := recover()

	if r != nil {
		fmt.Println("RECOVER", r)
	}
}

func forErrorHandling() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
