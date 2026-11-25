package main

import "fmt"

func getStringBytes() {

	var my_string string = "Hola Mundo"
	// Valor de UTF-8
	fmt.Println(my_string[0]) // -> 72

	var my_byte byte = my_string[0]
	fmt.Println(my_byte) // -> 72

	var s2 string = my_string[3:7]
	var s3 string = my_string[:7]
	fmt.Println("string 2: ", s2)
	fmt.Println("string 3: ", s3)

	// String --> Slices
	// Go representa los string como bytes
	var s string = "Hola Mundo"
	var bs []byte = []byte(s)
	var rs []rune = []rune(s)
	fmt.Println(bs)
	fmt.Println(rs)
}
