package main

import "fmt"

var a, b, c int                   // Variables del mismo tipo
var name, city = "Juan", "Madrid" // Variables del mismo tipo inferido
var x, y = 10, 3.1415             // Variables de diferentes tipos

// Usar en una sola línea, cuando se declaran varias variables del mismo tipo, se puede hacer de la siguiente manera:
// Ejemplo de uso cuando se declaran varias variables del mismo tipo
var (
	edad     int    = 30
	isActive bool   = true
	nickname string = "JJ"
)

func getMultipleVariables() {

	c, d := 10, 20
	full_name, ok := "Juan Perez", true

	c, d, e := 20, 30, 10

	fmt.Println("Variables a, b, c:", a, b, c)
	fmt.Println("Variables name, city:", name, city)
	fmt.Println("Variables x, y:", x, y)
	fmt.Println("Variables edad, isActive, nickname:", edad, isActive, nickname)
	fmt.Println("Variables c, d:", c, d)
	fmt.Println("Variables full_name, ok:", full_name, ok)
	fmt.Println("Variables c, d, e:", c, d, e)
}
