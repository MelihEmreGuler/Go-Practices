package main

import "fmt"

func swap(a, b int) (int, int) {
	return b, a
}

func main() {
	a, b := 2, 4
	fmt.Printf("a: %v b: %v\n", a, b)
	a, b = swap(a, b)

	fmt.Printf("a: %v b: %v", a, b)
}
