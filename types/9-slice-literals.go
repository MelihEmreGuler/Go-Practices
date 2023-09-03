package main

import "fmt"

type foo struct {
	i int
	b bool
}

func main() {

	s := []struct {
		i int
		b bool
	}{
		{1, true}, {2, false}, {3, true}, {4, false},
	}

	fmt.Println(s)

	array1 := []foo{
		{1, true}, {2, false}, {3, true}, {4, false},
	}

	fmt.Println(array1)

}
