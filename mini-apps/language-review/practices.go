package main

import "fmt"

var x int

type hotdog int

type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	lisenceToKill bool
}

type human interface {
	speak()
}

func (p person) speak() {
	fmt.Println(p.fname, p.lname, `says, "Good morning, James."`)
}

func (sa secretAgent) speak() {
	fmt.Println(sa.fname, sa.lname, `says, "Shaken, not stirred."`)
}

func saySomething(h human) {
	h.speak()
}

func main() {
	x = 7
	println(x)
	fmt.Printf("%T", x)
	xi := []int{2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(xi)

	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m["James"])

	var h hotdog
	h = 42
	fmt.Println(h)

	p1 := person{
		fname: "Miss",
		lname: "Moneypenny",
	}
	fmt.Println(p1)

	p1.speak()

	sa1 := secretAgent{
		person: person{
			fname: "James",
			lname: "Bond",
		},
		lisenceToKill: true,
	}
	fmt.Println(sa1)
	sa1.speak()

	sa1.person.speak() // This is how you can access the embedded type's method from the outer type.

	saySomething(p1)
	saySomething(sa1) // This is how you can pass a value of a type that implements an interface to a function that takes an interface as a parameter.
}
