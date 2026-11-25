package main

import "fmt"

func getValues() {

	// & => Es un operador que obtiene la dirección de memoria de la variable, lo que permite que la función Scanln pueda modificar el valor de la variable original.
	var name string
	var age int

	fmt.Println(name, age)
	fmt.Println(&name)
	fmt.Println(&age)

	fmt.Printf("Ingresa tu nombre: ")
	fmt.Scanln(&name)

	fmt.Printf("Ingresa tu edad: ")
	fmt.Scanln(&age)

	fmt.Printf("Hola %s, tienes %d años.\n", name, age)

	fmt.Println(name, age)
	fmt.Println(&name)
	fmt.Println(&age)
}
