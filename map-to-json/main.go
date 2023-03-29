package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	a := map[string]int{"a": 1, "b": 2}

	j, _ := json.Marshal(a) // marshal map to json

	println(string(j))

	//---------------------------------------------

	b := `{"a":1,"b":2}`

	var data map[string]int

	err := json.Unmarshal([]byte(b), &data) //unmarshal json to map

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(data)
	fmt.Println(data["a"])

	//---------------------------------------------

	c := `{"a":1,"b": [1,2,3]}` // json with array

	var data2 map[string]interface{}

	err2 := json.Unmarshal([]byte(c), &data2) //unmarshal json to map

	if err2 != nil {
		fmt.Println(err2)
	}

	fmt.Println(data2)
	fmt.Println(data2["a"])
	fmt.Println(data2["b"].([]interface{})[0], data2["b"].([]interface{})[1]) // type assertion
}
