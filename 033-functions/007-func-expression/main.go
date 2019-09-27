package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("My first func expression")
	}
	f()

	g := func(x int) {
		fmt.Println("The year big brother watching", x)
	}
	g(1984)
}
