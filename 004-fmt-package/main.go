package main

import (
	"fmt"
)

var y = 42

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)  // %T a Go-syntax representation of the type of the value
	fmt.Printf("%b\n", y)  // %b base 2
	fmt.Printf("%x\n", y)  // %x base 16, with upper-case letters for A-F
	fmt.Printf("%#x\n", y) // adds 0x at the begining of hexadecimal
	y = 911
	fmt.Printf("%#x\n", y)
	fmt.Printf("%#x\t%b\t%x\n", y, y, y)

	s := fmt.Sprintf("%#x\t%b\t%x\n", y, y, y)
	fmt.Println(s)
	fmt.Printf("%v", y) //the value in a default format

}
