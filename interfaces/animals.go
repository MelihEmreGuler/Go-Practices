package main

import "fmt"

type Animal interface {
	Says() string
	NumberOfLegs() int
}
type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func PrintInfo(a Animal) {
	fmt.Println("This animal says", a.Says(), "and has", a.NumberOfLegs(), "legs")
}

func (d *Dog) Says() string { //pointer receiver is best practice and better performance
	return "Woof"
}

func (d *Dog) NumberOfLegs() int { //pointer receiver is best practice and better performance
	return 4
}

func (g *Gorilla) Says() string {
	return "Ugh"
}

func (g *Gorilla) NumberOfLegs() int {
	return 2
}
