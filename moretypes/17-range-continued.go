package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow { //index i saymaya ihtiyac duymuyorsak bos tanimlayici kullanabiliriz
		fmt.Printf("%d\n", value)
	}
	
}
