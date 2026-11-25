package main

import "fmt"

const Pi float64 = 3.141519

func typeConstants() {
	// Constante NO TIPADA es una constante que no tiene un tipo de dato específico y puede ser utilizada en diferentes contextos.
	var number float64 = Pi

	fmt.Println("Valor de Pi:", number)
}
