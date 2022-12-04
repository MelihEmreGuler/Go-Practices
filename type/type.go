package main

import "fmt"

var a int

type hotdog int

var b hotdog

func main() {
	a = 42
	b = 43
	fmt.Printf("a: %T\n", a)
	fmt.Printf("b: %T\n", b)

	// a = b (different types)
	
}
