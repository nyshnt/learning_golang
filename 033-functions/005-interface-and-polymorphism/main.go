package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

type human interface {
	speak()
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last)
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, " - the person speak")
}

func bar(h human) {
	fmt.Println("I was passed into bar", h)
}

func main() {
	sa1 := secretAgent{
		person: person{
			"james",
			"Bond",
		},
		ltk: true,
	}
	sa2 := secretAgent{
		person: person{
			"Miss",
			"MOneypenny",
		},
		ltk: true,
	}

	p1 := person{
		first: "Dr.",
		last:  "Yes",
	}

	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()

	fmt.Println(p1)

	bar(sa1)
	bar(sa2)
	bar(p1)
}
