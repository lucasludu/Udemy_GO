package main

import "fmt"

func getSlicesOfSlices() {

	// Slice de Slice -> Son referencias a un arreglo subyacente, por lo que si se modifica un slice, se modifica el arreglo subyacente y por lo tanto los demás slices que referencian ese mismo arreglo.

	x := []string{"a", "b", "c", "d", "e"}

	// Slice de Slice -> Se puede crear un slice a partir de otro slice, lo que permite crear sub-slices que comparten el mismo arreglo subyacente.

	y := x[0:2] // -> ["a", "b"]
	z := x[1:3] // -> ["b", "c"]
	d := x[1:4] // -> ["b", "c", "d"]
	// e := x[:]   // -> ["a", "b", "c", "d", "e"]
	e := make([]string, 5)
	copy_e := copy(e, x)

	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
	fmt.Println("d:", d)
	fmt.Println("e:", e, copy_e)

	// Esto es una copia de la dirección de memoria del arreglo subyacente, por lo que si se modifica un slice, se modifica el arreglo subyacente y por lo tanto los demás slices que referencian ese mismo arreglo.
	e[0] = "x"
	fmt.Println("x:", x)
	fmt.Println("e:", e)
}
