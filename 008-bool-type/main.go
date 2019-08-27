package main

import "fmt"

var x bool

func main() {
	fmt.Println(x)
	x = true
	fmt.Println(x)

	a := 14
	b := 15
	fmt.Println(a != b)
}
