package main

import "fmt"

func main() {
	/*
		colors := map[string]string{
			"red":   "#ff0000",
			"green": "#00ff00",
		}
	*/

	//var colors map[string]string
	/*
		colors := make(map[string]string)

		colors["white"] = "#ffffff" //map' imizce yeni bir deger ekledik

		fmt.Println(colors)

		delete(colors, "white") //map' imizde key ini girdigimiz degeri sildik

		fmt.Println(colors)

	*/

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
		"white": "#ffffff",
	}
	printMap(colors)

}
func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
