package main

import "fmt"

func getBasicTypes() {
	var a int8 = 127
	var b float32 = 3.14
	var c string = "Hola Mundo"
	var d bool = true
	var e rune = 'A'
	var f complex128 = 2 + 3i

	// Imprimir los valores y sus tipos
	fmt.Printf("a= %d (%T)\n", a, a)
	fmt.Printf("b= %f (%T)\n", b, b)
	fmt.Printf("c= %s (%T)\n", c, c)
	fmt.Printf("d= %t (%T)\n", d, d)
	fmt.Printf("e= %c (%T)\n", e, e)
	fmt.Printf("f= %v (%T)\n", f, f)
}
