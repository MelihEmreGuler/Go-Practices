package main

import "fmt"

func main() {
	pow := []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow { //range kullanımını (in pow)(foreach)
		fmt.Printf("2**%v = %v \n", i, v)
	}
}
