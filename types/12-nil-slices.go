package main

import "fmt"

func main() {
	var slice []int
	fmt.Printf("len: %v, cap: %v, slice: %v\n", len(slice), cap(slice), slice)

	if slice == nil {
		fmt.Println(slice, "is", nil)
	} //The zero value of a slice is nil.
}
