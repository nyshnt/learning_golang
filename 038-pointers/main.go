package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a)
	nn
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a) // & gives you the address

	b := &a
	fmt.Println(b)
	fmt.Println(*b)
	fmt.Println(*&a) // * gives you the value stored at an address when you have the address

	*b = 43
	fmt.Println(a)
}
