package main

import "fmt"

func getMapInit() {

	// Map Literal
	/*
		teams := map[string][]string{
			"Carlito Figuero CF:": []string{"Carlito", "Fernando", "Anchelo"},
			"Los Inmaduros":       []string{"Bilhere", "Zunino", "Genga"},
			"Los Instructores":    []string{"Gabriel", "Teddy", "Ricardo"},
		}

		fmt.Println(teams["Los Instructores"][0])
	*/

	ages := make(map[int][]string, 10)
	fmt.Println(ages)

	// Todos los que no se puedan comparar, es decir, usar == o !=

}
