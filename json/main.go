package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//write struct from a json
	myJson := `[
	{
		"first_name": "Clark",
		"last_name": "Kent",
		"hair_color": "black",
		"has_dog": true
	},
	{
		"first_name": "Bruce",
		"last_name": "Wayne",
		"hair_color": "black",
		"has_dog": false
	}
]`
	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJson), &unmarshalled) //unmarshal json to map
	if err != nil {
		fmt.Println("Error unmarshalling json", err)
	}

	fmt.Printf("unmarshalled: %v\n", unmarshalled)

	//write json from a struct

	var mySlice []Person

	m1 := Person{"Wally", "West", "Red", false}
	mySlice = append(mySlice, m1)
	m2 := Person{"Diana", "Prince", "Black", true}
	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, "", "	") //indent with tab, newJson is a slice of bytes
	if err != nil {
		fmt.Println("error marshalling", err)
	}

	fmt.Println(string(newJson)) //convert slice of bytes to string
}
