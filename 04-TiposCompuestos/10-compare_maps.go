package main

import (
	"fmt"
	"maps"
)

func getCompareMaps() {

	my_map_a := map[string]int{
		"Hola":  1,
		"Mundo": 2,
	}

	my_map_b := map[string]int{
		"Hola":  1,
		"Mundo": 2,
	}

	// fmt.Println(my_map_a, my_map_b)
	fmt.Println(maps.Equal(my_map_a, my_map_b))
}
