package main

import "fmt"

func main() {
	myNumber := 5
	i := &myNumber // var i *int = &myNumber

	fmt.Println(myNumber, i)

	*i = 6

	fmt.Println(myNumber, i)

	// newPtr := i + 1 // go' da pointer aritmatigi yoktur bu satir hata verir
}
