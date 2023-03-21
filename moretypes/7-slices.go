package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13} // array of 6 integers

	var slicedPrimes []int     // slice of 3 integers (it's allocated empty integers by default) [0 0 0]
	slicedPrimes = primes[1:4] // slice of primes from index 1 to 4 (not included)

	fmt.Println(slicedPrimes) //This selects a half-open range which includes the first element, but excludes the last one. [3 5 7]

	slicedPrimes = append(slicedPrimes, 11, 17)
	fmt.Println(slicedPrimes) // [3 5 7 11 17]

}
