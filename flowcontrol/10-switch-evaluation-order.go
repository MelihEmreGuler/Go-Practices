package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("When is Saturday?")

	today := time.Now().Weekday()

	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	a := 2

	switch a { //fallthrough is used to execute the next case
	case 1:
		fmt.Println("a is 1")
		fallthrough
	case 2:
		fmt.Println("a is 2")
		fallthrough
	case 3:
		fmt.Println("a is 3")
		fallthrough
	case 4:
		fmt.Println("a is 4")
		fallthrough
	default:
		fmt.Println("a is not 1,2,3,4")
	}

	switch a := 10; { //switch with no condition is same as switch true
	case a > 0:
		fmt.Println("a is greater than 0")
	case a < 0:
		fmt.Println("a is less than 0")
	default:
		fmt.Println("a is 0")
	}
}
