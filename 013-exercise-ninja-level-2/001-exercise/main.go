// Write a program that prints a number in decimal, binary, and hex
package main

import "fmt"

func main() {
	x := 14
	fmt.Printf("%d\n", x)
	fmt.Printf("%b\n", x)
	fmt.Printf("%#x", x)
}
