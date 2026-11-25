package main

import "fmt"

func getOperadoresComparacion() {

	age := 18
	isVIP := false

	fmt.Println(age >= 18)     // true
	fmt.Println(isVIP == true) // false
	fmt.Println(isVIP != true) // true
}
