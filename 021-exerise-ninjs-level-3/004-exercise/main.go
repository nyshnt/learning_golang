// Create a for loop using this syntax
// for { }
// Have it print out the years you have been alive.

package main

import "fmt"

func main() {
	bd := 1993
	for {
		if bd > 2019 {
			break
		}
		fmt.Println(bd)
		bd++
	}
}
