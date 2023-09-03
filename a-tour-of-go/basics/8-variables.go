package main

import "fmt"

var (
	c, python, java bool

	MaxInt int = 1<<32 - 1

	myString string
)

func main() {
	fmt.Println(c, python, java) //default false atanir

	fmt.Printf("int: %v\n", MaxInt)

	fmt.Printf("string: %q", myString) //stringler icin %q da kullanilabilir
}
