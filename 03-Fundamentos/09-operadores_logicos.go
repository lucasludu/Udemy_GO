package main

import "fmt"

func getOperadoresLogicos() {

	a := true
	b := false

	// and (&&) --> Todos los valores deben ser true para que el resultado sea true
	fmt.Println("a && b = ", a && b)

	// or (||) --> Al menos un valor debe ser true para que el resultado sea true
	fmt.Println("a || b = ", a || b)

	// not (!) --> Invierte el valor de la variable
	fmt.Println("!a = ", !a)
	fmt.Println("!b = ", !b)

}
