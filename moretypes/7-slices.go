package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var slicedPrimes []int = primes[1:4]

	fmt.Println(slicedPrimes) //This selects a half-open range which includes the first element, but excludes the last one.
}
