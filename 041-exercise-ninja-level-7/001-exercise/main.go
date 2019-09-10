// create a value and assign it to a variable
// print the address of that value

package main

import "fmt"

func main() {
	a := 14
	fmt.Println("Value 	:", &a)
	fmt.Println("Address :", &a)
}
