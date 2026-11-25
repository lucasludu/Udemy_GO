package main

import "fmt"

func zeroValue() {
	var i int
	var s string
	var ok bool
	var nums []int

	fmt.Printf("i= %d\n", i)

	fmt.Printf("s= %s\n", s)
	fmt.Printf("s= %q\n", s)

	fmt.Printf("ok= %t\n", ok)
	fmt.Printf("ok= %v\n", ok)

	fmt.Printf("nums= %v\n", nums)
}
