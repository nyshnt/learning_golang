package main

import (
	"fmt"
)

var a int

type daemon int // created own type

var b daemon

func main() {
	a = 14
	b = 15
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	a = int(b) // conversion
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
