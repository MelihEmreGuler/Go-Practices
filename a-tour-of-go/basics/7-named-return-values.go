package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 2
	y = x + sum

	return //x ve y yi return edecek
}

func main() {
	fmt.Println(split(4))
}
