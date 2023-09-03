package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	if a := 1; a > 0 { //different type of if
		fmt.Println("a is positive")
	} else {
		fmt.Println("a is negative")
	}
}

/*
cikti
27 >= 20
9 20
once fonksiyonun içerisindeki islemleri gerceklestiriyor en sonunda fonksiyonun donus degerini gerceklestiriyor

*/
