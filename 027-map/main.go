package main

import "fmt"

func main() {
	m := map[string]int{
		"Tom":   3,
		"Jerry": 2,
	}
	fmt.Println(m)
	fmt.Println(m["Tom"])
	fmt.Println(m["Spike"])

	v, ok := m["Spike"]
	fmt.Println(v)
	fmt.Println(ok)

	m["Butch Cat"] = 4 // Adding an element in Map
	if v, ok := m["Jerry"]; ok {
		fmt.Println("This is IF Print", v)
	}

	for i, v := range m {
		fmt.Println(i, v)
	}

	xi := []int{4, 5, 6, 7, 8, 9}

	for i, v := range xi {
		fmt.Println(i, v)
	}

	// Delete an entry form map

	delete(m, "Tom")
	fmt.Println("After delete ", m)

	if v, ok := m["Jerry"]; ok {
		fmt.Println("value:", v)
		delete(m, "Jerry")
	}
	fmt.Println(m)
}
