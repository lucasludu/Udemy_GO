package main

import "fmt"

var age int

func getDeclarations() {
	age = 34
	name := "John"

	fmt.Println("Age:", age)
	fmt.Println("Name:", name)
}
