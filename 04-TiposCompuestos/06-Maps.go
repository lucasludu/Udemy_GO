package main

import "fmt"

func getMaps() {

	// Map[KeyType]ValueType
	// var nilMap map[string]int
	// nilMap["Hola"] = 1 --> Panic -> Es un error

	totalWins := map[string]int{}
	// fmt.Println(totalWins == nil) // -> Da un "False"
	totalWins["Los domadores"] = 1
	totalWins["Los sedentarios"] = 2
	totalWins["Fernando Robles FR"] = 5

	fmt.Println(totalWins)
	fmt.Println(totalWins["Fernando Robles FR"])

}
