package main

import "fmt"

func main() {
	x := make([]int, 10, 100)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x[0] = 1
	x[9] = 2

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}
