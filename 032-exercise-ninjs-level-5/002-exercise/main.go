// Take the code from the previous exercise,
// then store the values of type person in a map with the key of last name.
// Access each value in the map.
// Print out the values, ranging over the slice.

package main

import "fmt"

type person struct {
	first       string
	last        string
	favFlavours []string
}

func main() {
	p1 := person{
		first:       "Tony",
		last:        "Stark",
		favFlavours: []string{"chocolate", "vanilla"},
	}
	p2 := person{
		first:       "James",
		last:        "Bond",
		favFlavours: []string{"butterscotch", "strawberry"},
	}

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, v := range m {
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, val := range v.favFlavours {
			fmt.Println(i, val)
		}
		fmt.Println("---------")
	}
}
