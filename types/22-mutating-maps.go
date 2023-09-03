package main

import "fmt"

func main() {

	var m map[string]int
	m = make(map[string]int)
	m["key1"] = 1

	map1 := make(map[string]int)
	map1["key1"] = 1

	fmt.Println(m, map1)

	map2 := map[string]int{"first": 1, "second": 2, "third": 3}

	fmt.Println(map2["second"])

	delete(map2, "second") // delete the key and value "second:2" from the map

	fmt.Println(map2) // map[first:1 third:3]

	fmt.Println(map2["second"]) // 0

	value, b := map2["third"] // b is a boolean value, if the key exists in the map, b will be true, otherwise it will be false
	fmt.Println(value, b)      // 3 true
}
