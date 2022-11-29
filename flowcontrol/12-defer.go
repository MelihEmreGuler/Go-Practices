package main

import "fmt"

func main() {
	defer fmt.Println("First statement")
	fmt.Println("Second statement")
	fmt.Println("Third statement")
	fmt.Println("Fourth statement")

}
