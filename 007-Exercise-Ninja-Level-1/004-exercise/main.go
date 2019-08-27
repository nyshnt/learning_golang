package main

import "fmt"

type super int

var x super

func main() {

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)

}
