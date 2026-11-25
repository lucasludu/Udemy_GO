package main

import (
	"fmt"
	"slices"
)

func getCompareSlice() {

	x := []int{1, 2, 3, 4, 5}
	y := []int{1, 2, 3, 4, 5}
	z := []int{1, 2, 3, 4, 5, 6}
	// s := []string{"a", "b", "c"}

	fmt.Println(slices.Equal(x, y))
	fmt.Println(slices.Equal(x, z))

	// NO podemos comparar slices de diferentes tipos, ya que cada tipo de slice tiene su propia representación interna y estructura de datos. Comparar slices de diferentes tipos no tendría sentido, ya que podrían tener diferentes longitudes, capacidades y elementos subyacentes, lo que haría que la comparación fuera inválida.
	// fmt.Println(slices.Equal(x, s))

}
