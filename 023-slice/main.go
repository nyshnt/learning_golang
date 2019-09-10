package main

import "fmt"

func main() {
	// // COMPOSITE LITERAL
	// x := []int{4, 5, 7, 8, 42}
	// fmt.Println("x\n", x)
	// x1 := x[2:4]
	// fmt.Println("x1\n", x1)
	// x2 := x[:3]
	// fmt.Println("x2\n", x2)
	// x1[0] = 10
	// fmt.Println("x1\n", x1)
	// fmt.Println("x\n", x)

	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set: ", s)
	fmt.Println("get: ", s[2])

	fmt.Println("len: ", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd :", s)

	fmt.Println("New len: ", len(s))

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy :", c)

	l := s[2:5]
	fmt.Println("sl1: ", l)

	l = s[:5]
	fmt.Println("sl2 :", l)

	l = s[2:]
	fmt.Println("sl3 :", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
		fmt.Println("2d: ", twoD)
	}

}

// a SLICE literal allows to group together VALUES of the same Type
