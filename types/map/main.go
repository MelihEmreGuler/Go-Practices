package main

import "fmt"

type person struct {
	fname string
	lname string
}

func main() {

	colors1 := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
	}

	//var colors map[string]string //map' in degeri nil'dir, nil map' e deger ekleyemeyiz

	colors2 := make(map[string]string)
	fmt.Println("colors2:", colors2)

	colors1["white"] = "#ffffff" //map' imizce yeni bir deger ekledik
	fmt.Println("colors1:", colors1)

	delete(colors1, "white") //map' imizde key ini girdigimiz degeri sildik

	fmt.Println("colors1", colors1)

	colors3 := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
		"white": "#ffffff",
	}
	printMap(colors3)

	myMap := make(map[string]person)

	myMap["me"] = person{fname: "Melih Emre", lname: "Guler"}

	myMap["me"].printPerson()
}
func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}

func (p person) printPerson() person {
	fmt.Printf("first name: %s, last name: %s\n", p.fname, p.lname)
	return p
}
