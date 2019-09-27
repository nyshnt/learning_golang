// Create your own type “person” which will have an underlying type of “struct”
// so that it can store the following data:
// first name
// last name
// favorite ice cream flavors
// Create two VALUES of TYPE person. Print out the values,
// ranging over the elements in the slice which stores the favorite flavors.

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
		favFlavours: []string{"chocolate", "vanila"},
	}
	p2 := person{
		first:       "Tony",
		last:        "Stark",
		favFlavours: []string{"chocolate", "vanila"},
	}

	fmt.Println(p1.first)
	fmt.Println(p1.last)
	for i, v := range p1.favFlavours {
		fmt.Println(i, v)
	}

	fmt.Println(p2.first)
	fmt.Println(p2.last)
	for i, v := range p2.favFlavours {
		fmt.Println(i, v)
	}

}
