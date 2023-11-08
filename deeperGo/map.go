package main

import "fmt"

func maping() {
	fmt.Println("map is start from here")
	arrmap := map[string]int{"ajay": 55, "vijay": 66, "sanjay": 45, "dhananjay": 55, "jay": 44}
	fmt.Println(arrmap["jay"])
	arrmap["Ridesh"] = 22
	fmt.Println(arrmap)

	delete(arrmap,"jay")
	fmt.Println(arrmap)
	fmt.Println(arrmap["sanjay"])


	for key, val := range arrmap{

		if key == "Ridesh" {
			fmt.Println("Welcome Ridesh Ji")
			arrmap[key]= 88
		}

		if val > 50 {
			fmt.Println("Namste " + key + " Ji ")
		}
	}
	fmt.Println(arrmap)

}
