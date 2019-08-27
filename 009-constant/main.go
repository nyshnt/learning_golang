package main

import "fmt"

const (
	// untyped constants
	a = 14
	b = 15.67
	c = "Iron Man"
)

const (
	// typed constants
	x int     = 15
	y float64 = 16.78
	z string  = "Thor"
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", z)
}
