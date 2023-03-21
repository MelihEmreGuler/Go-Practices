package main

import "fmt"

func main() {
	pow := []int{0, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow { //usage of range
		//v := pow[i] // assign the element value to v
		fmt.Printf("2*%v = %v \n", i, v)
	}

	for _, v := range pow { // usage of range with _
		fmt.Printf("%v \n", v)
	}

	for i := range pow { // usage of range with only index
		fmt.Printf("%v \n", i) // 0 1 2 3 4 5 6 7
	}

	// range on map iterates over key/value pairs.
	var m map[string]int
	m = make(map[string]int)
	m["key1"] = 1
	m["key2"] = 2
	m["key3"] = 3
	for k, v := range m {
		fmt.Printf("%v: %v \n", k, v) // key1: 1 key2: 2 key3: 3
	}

}
