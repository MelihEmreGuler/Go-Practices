package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i) // 0 1 2 3 4 5 6 7 8 9 //last in first out
	}

	fmt.Println("done")
}
