package main

import "fmt"

func main() {
	foo()
	bar("James")
	s1 := woo("Moneypenny")
	fmt.Println(s1)
	x, y := mouse("Ian", "Fleming")
	fmt.Println(x)
	fmt.Println(y)
}

// func (r receiver) identitifier(parameters) (return(s)) { ... }

func foo() {
	fmt.Println("Hello from foo")
}

// Everythig in Go is PASS BY VALUE
func bar(s string) {
	fmt.Println("Hello", s)
}

func woo(s string) string {
	return fmt.Sprint("Hello from woo, ", s)
}

func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ", ln, `, Says "Hello"`)
	b := false
	return a, b
}
