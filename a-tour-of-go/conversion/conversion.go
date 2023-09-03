package main

import "fmt"

var a int

type hotdog int

var b hotdog

func main() {
	a = 42
	b = 43

	a = int(b) //converison

	fmt.Printf("a: %T, b: %T\n", a, b)
	//a: int, b: main.hotdog (no changes)

}
