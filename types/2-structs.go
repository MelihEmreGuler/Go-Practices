package main

import "fmt"

type Coordinates struct {
	X int
	Y int
}

func main() {

	var coordinates1 Coordinates = Coordinates{
		X: 5,
		Y: 4,
	}
	fmt.Println(coordinates1)

	fmt.Println(Coordinates{X: 2, Y: 3}) //kendi toString methodu var.

	coordinates2 := Coordinates{11, 22}
	fmt.Println(coordinates2.X)

	coordinates3 := Coordinates{X: 33} //value of Y is 0
	fmt.Println(coordinates3)

}
