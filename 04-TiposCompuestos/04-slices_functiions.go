package main

import "fmt"

func getSliceInfo() {

	/*
		var my_slice []int
		fmt.Println(my_slice == nil)
	*/

	/* PARTE 1
	my_slice := []int{1, 2, 3, 4, 5}

	// Lend
	fmt.Println("Length of slice:", len(my_slice))
	// Capacity
	fmt.Println("Capacity of slice:", cap(my_slice))

	fmt.Println("------------------------------------------")

	// Append
	my_slice = append(my_slice, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("Slice after append:", my_slice)
	fmt.Println("Length of slice after append:", len(my_slice))
	fmt.Println("Capacity of slice after append:", cap(my_slice))

	fmt.Println("------------------------------------------")
	// Len y Capacity -> La diferencia entre len y cap es que len devuelve el número de elementos en el slice, mientras que cap devuelve la capacidad del slice, es decir, el número máximo de elementos que puede contener antes de necesitar una reasignación de memoria.

	fmt.Println(my_slice, len(my_slice), cap(my_slice))
	fmt.Println(len(my_slice) == cap(my_slice))
	*/

	/* PARTE 2
	 */

	make_slice := make([]int, 5, 10)
	fmt.Println("Slice created with make:", make_slice)
	fmt.Println("Length of slice:", len(make_slice))
	fmt.Println("Capacity of slice:", cap(make_slice))

	fmt.Println("------------------------------------------")
	
	// Si tengo un slice con capacidad 10 y le agrego 6 elementos, la capacidad se duplicará a 20. Esto se debe a que cuando se alcanza la capacidad máxima de un slice, Go crea un nuevo slice con el doble de capacidad y copia los elementos del slice original al nuevo slice.
	
	make_slice = append(make_slice, 1, 2, 3, 4, 5, 10)
	fmt.Println("Slice after append:", make_slice)
	fmt.Println("Length of slice after append:", len(make_slice))
	fmt.Println("Capacity of slice after append:", cap(make_slice))

	fmt.Println("------------------------------------------")
	// Vaciar un slice
	// make_slice = make_slice[0:0] // --> 0, 20
	clear(make_slice) // --> 11 , 20
	fmt.Println("Slice after clearing:", make_slice)
	fmt.Println("Length of slice after clearing:", len(make_slice))
	fmt.Println("Capacity of slice after clearing:", cap(make_slice))
}
