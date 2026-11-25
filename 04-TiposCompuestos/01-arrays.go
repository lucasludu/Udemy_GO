package main

import "fmt"

func getArrays() {
	// var number_list [3]int

	// Array Literal --> van a tener valores
	// var number_list = [3]int{10, 20, 30}
	// Array Literal Simplificado --> van a tener valores
	var number_list = [...]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Println(number_list)

	number_list[0] = 5
	fmt.Println(number_list)

	fmt.Println(number_list[2])
	// fmt.Println(number_list[4]) // Esto va a dar error porque no existe el índice 4

	// Para obtener la longitud de un array se utiliza la función len()
	fmt.Println(len(number_list))
}
