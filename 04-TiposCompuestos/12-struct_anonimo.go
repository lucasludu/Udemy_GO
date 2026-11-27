package main

import "fmt"

func getStructAnonimo() {

	// Struct Anonimo
	var person struct {
		name string
		age  int
		pet  string
	}

	person.name = "Carlito"
	person.age = 40
	person.pet = "Perro"

	fmt.Println(person)

	pet := struct {
		name string
		kind string
	}{
		name: "Dante",
		kind: "Perro",
	}

	fmt.Println(pet)

}
