package main

import "fmt"

func getCommaOkIdiom() {

	my_map := map[string]int{
		"Hola":  1,
		"Mundo": 2,
	}

	value, ok := my_map["Hola"]
	// value, ok := my_map["Chanchito"]

	/*
		fmt.Println(my_map)
		fmt.Println(my_map["Hola"])
		fmt.Println(my_map["Chanchita"])
	*/

	fmt.Println(value, ok)
}
