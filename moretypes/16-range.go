package main

import "fmt"

func main() {
	pow := []int{0, 2, 4, 8, 16, 32, 64, 128}
	animals := []string{"dog", "cat", "bird", "fish", "cow", "horse", "sheep", "goat"}

	var i = 0
	for range animals {
		i++
	}
	fmt.Printf("i: %v \n", i)

	for i, animal := range animals {
		fmt.Printf("i: %v animal: %s \n", i, animal)
	}

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

	animalsMap := map[string]string{
		"dog": "dirty",
		"cat": "orange",
	}
	animalsMap["bird"] = "bidik"

	for animalType, animalName := range animalsMap {
		fmt.Printf("%v.: %s \n", animalType, animalName)
	}

	type User struct {
		Name     string
		LastName string
		EMail    string
		Age      int
	}
	var users []User

	users = append(users, User{"melih", "guler", "melih@guler.com", 22})
	users = append(users, User{"emre", "guren", "emre@guren.com", 20})
	users = append(users, User{"mert", "senel", "mert@senel.com", 6})

	for i, user := range users {
		fmt.Printf("%v. user name: %v, last name: %v, e mail: %v, age: %v \n", i, user.Name, user.LastName, user.EMail, user.Age)
	}
}
