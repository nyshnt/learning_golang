package main

import "fmt"

func main() {
	x := 131
	if x == 14 {
		fmt.Println("our value was 14")
	} else if x == 15 {
		fmt.Println("our value was 15")
	} else if x == 16 {
		fmt.Println("our value was 16")
	} else {
		fmt.Println("our value was not 14, 15, or 16")
	}
}
