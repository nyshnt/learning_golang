package main

import "fmt"

func main() {
	// COMPOSITE LITERAL
	x := []int{4, 5, 7, 8, 42}
	fmt.Println("x\n", x)
	x1 := x[2:4]
	fmt.Println("x1\n", x1)
	x2 := x[:3]
	fmt.Println("x2\n", x2)
	x1[0] = 10
	fmt.Println("x1\n", x1)
	fmt.Println("x\n", x)
}

// a SLICE literal allows to group together VALUES of the same Type
