package main

import "fmt"

func loops() {
	fmt.Println("Now start loops from here :- ")
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		if i == 8 {
			break
		}
		fmt.Print(i, "\n")
	}
	var n int
	fmt.Println("Get table of given Input ")
	fmt.Scanf("%d",&n)
	for i := 1; i <= 10; i ++ {
		fmt.Println(n,"*",i,":- ",n*i)
	}
}
