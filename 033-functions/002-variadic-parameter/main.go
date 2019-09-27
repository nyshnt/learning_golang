package main

import "fmt"

func main() {
	x := sum(2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("The tota is ", x)
}

func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("For item in index position,", i, "We are now adding,", v, "to the total which is now", sum)
	}
	return sum
}

// func (r receiver) identifier(parameter(s)) (return(s)) {code}
