package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println("i:", i)
	}

	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println("when the sum smaller than 1000 repeat. sum:", sum)

}
