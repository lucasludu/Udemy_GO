package main

import "fmt"

func getSlice() {

	/*
		var array1 = [3]int{1, 2, 3}
		var array2 = [4]int{1, 2, 3}
	*/

	// El slice es un tipo de dato compuesto que permite almacenar una colección de elementos del mismo tipo, pero a diferencia de los arrays, los slices son dinámicos y pueden crecer o reducir su tamaño según sea necesario. Los slices se definen utilizando corchetes vacíos [] y no requieren especificar un tamaño fijo.
	var my_slice = []int{1, 2, 3, 4, 5}
	my_slice = append(my_slice, 6, 7, 8, 9, 10)
	fmt.Println(my_slice)

	my_slice[1] = 23
	fmt.Println(my_slice)
	fmt.Println(my_slice[1])

	fmt.Println("--------------------------------------")
	fmt.Println("--------------------------------------")

	var zero_slice []int
	// Un slice vacío es un slice que no contiene elementos. Se puede crear utilizando la función make() o simplemente declarando un slice sin inicializarlo. Un slice vacío tiene una longitud de 0 y una capacidad de 0, lo que significa que no puede almacenar ningún elemento hasta que se le agreguen elementos mediante la función append().
	// nil es un valor especial en Go que representa la ausencia de un valor o una referencia nula. Cuando se compara un slice con nil, se está verificando si el slice no ha sido inicializado o si no tiene elementos asignados. Un slice vacío no es igual a nil, ya que aunque no tenga elementos, sigue siendo un slice válido y puede ser utilizado para almacenar elementos en el futuro.
	fmt.Println(zero_slice)
	fmt.Println(zero_slice == nil)
	fmt.Println(len(zero_slice))
	fmt.Println(cap(zero_slice))

	fmt.Println("--------------------------------------")
	fmt.Println("--------------------------------------")

	// var other_zero_slice []int
	// Un slice vacío es un slice que no contiene elementos. Se puede crear utilizando la función make() o simplemente declarando un slice sin inicializarlo. Un slice vacío tiene una longitud de 0 y una capacidad de 0, lo que significa que no puede almacenar ningún elemento hasta que se le agreguen elementos mediante la función append().

	// No podemos comparar dos slices directamente utilizando el operador de igualdad (==) en Go. Esto se debe a que los slices son estructuras de datos que contienen un puntero a un array subyacente, su longitud y su capacidad. Comparar dos slices directamente no tiene sentido, ya que podrían apuntar a diferentes arrays subyacentes o tener diferentes longitudes y capacidades, incluso si contienen los mismos elementos.
	// --> fmt.Println(zero_slice == other_zero_slice)
	// // Esto generará un error de compilación: invalid operation: zero_slice == other_zero_slice (slice can only be compared to nil)

}
