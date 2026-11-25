package main

import "fmt"

// Declaración de constantes
const appInterno = "Curso de Go"

// Declaración de constantes con mayúscula inicial para exportarlas
const MaxUsers = 1000

func constants() {
	// La declaración de las variables tienen que estar dentro de una función, no se pueden declarar variables fuera de una función
	// Tienen que estar lo más cerca posible de donde se van a utilizar, para que el código sea más legible y fácil de mantener

	// Declaración de variables
	var name string = "Lucas"

	// Declaración de variable con inferencia de tipo
	lastname := "Gonzalez"

	fmt.Printf("Nombre: %s %s %T\n", name, lastname, lastname)
	fmt.Println(appInterno, MaxUsers)
}
