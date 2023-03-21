package main

import "fmt"

func main() {
	var a [2]string // array of 2 strings (it's allocated empty strings by default) [  ]

	var b [4]bool

	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a[0], a[1], a) // Hello World [Hello World]
	fmt.Println(b)             // [false false false false]

	fmt.Println("len(a):", len(a))
	fmt.Println("len(b):", len(b))

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fibonacci := [...]int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55} // [...] means that the compiler will count the number of elements in the array

	fmt.Println(primes)
	fmt.Println(fibonacci)
}
