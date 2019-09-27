// Using a COMPOSITE LITERAL:
// create an ARRAY which holds 5 VALUES of TYPE int
// assign VALUES to each index position.
// Range over the array and print the values out.
// Using format printing
// print out the TYPE of the array

package main

import "fmt"

func main() {
	x := [5]int{10, 11, 12, 13, 14}

	for i, v := range x {
		fmt.Println(i, v)
	}

	fmt.Printf("%T", x)

}
