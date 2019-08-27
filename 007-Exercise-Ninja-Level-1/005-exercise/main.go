package main

import "fmt"

type super int

var x super
var y int

func main() {

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 14
	fmt.Println(x)
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

}
