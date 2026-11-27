package main

import "fmt"

func getStruct() {

	// --> Tipo Persona
	type person struct {
		name string
		age  int
		pet  string
	}

	// carlitos := person{}

	// Struct Literal
	// fernando := person{}

	carlitos := person{
		"Carlos",
		35,
		"Perro",
	}

	fernando := person{
		age:  40,
		pet:  "Perro",
		name: "Fernando",
	}

	fmt.Println(carlitos, fernando)
	fernando.name = "Fernandinho"
	fmt.Println(carlitos.name, fernando.name)

}
