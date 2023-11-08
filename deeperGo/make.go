package main

import "fmt"

func makes() {
	fmt.Println("make start from here")
	// make is use to declare size of slice and define datatype for map once map

	myslice := make([]int, 5, 10)
	myslice = append(myslice, 5, 6, 4)
	myslice[1] = 55
	fmt.Println(len(myslice))
	fmt.Println(myslice)

	mymap := make(map[int]int)
	mymap[2] = 8
	mymap[4] = 5
	fmt.Println(mymap)	
}
