package main

import (
	"errors"
	"fmt"
)

func main() {

	result, err := divide(100.0, 10.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("result of division is %v\n", result)
}

func divide(x, y float32) (float32, error) {
	var result float32

	if y == 0 {
		return result, errors.New("cannot divide by zero")
	}

	result = x / y
	return result, nil //no error
}
