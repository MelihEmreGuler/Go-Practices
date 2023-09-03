package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	vertex1 := Vertex{X: 5, Y: 6}
	/*
		vertex2 := vertex1

		fmt.Printf("vertex1 X: %v, Y: %v\nptr X: %v Y: %v\n", vertex1.X, vertex1.Y, vertex2.X, vertex2.Y)

		vertex2.X = 7
		vertex2.Y = 8

		fmt.Printf("vertex1 X: %v, Y: %v\nptr X: %v Y: %v \n", vertex1.X, vertex1.Y, vertex2.X, vertex2.Y)
	*/

	ptr := &vertex1

	// *ptr.X = 4 aslında bu sekilde bir sntax olması lazım ama daha kolay kullanim için go struc larda x kullanımına gerek duymuyor
	ptr.X = 4
	fmt.Printf("vertex1 X: %v, Y: %v\nptr X: %v Y: %v \n", vertex1.X, vertex1.Y, ptr.X, ptr.Y)

	fmt.Println(ptr) // basinda & ile bir toString methodu var

	fmt.Println(v1, p, v2, v3)

}
