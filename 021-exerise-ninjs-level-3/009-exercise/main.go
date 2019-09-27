// Create a program that uses a switch statement with the switch expression specified
// as a variable of TYPE string with the IDENTIFIER “favSport”.

package main

import "fmt"

func main() {
	favSport := "swimming"

	switch favSport {
	case "cricket":
		fmt.Println("Go to the ground.")
	case "hiking":
		fmt.Println("Go to the hills.")
	case "swimming":
		fmt.Println("Go to the pool.")
	}
}
