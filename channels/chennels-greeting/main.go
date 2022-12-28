package main

import (
	"fmt"
)

func main() {
	c := make(chan string)
	go assignGreeting("Hello", c)

	fmt.Println(<-c)
}

func assignGreeting(message string, c chan string) {
	c <- message

}

//yanlis ornekler;

/* package main

import "fmt"

func main() {
	c := make(chan string)
	for i := 0; i < 4; i++ {
		go printString("Hello there!", c)
	}

	for s := range c {
		fmt.Println(s)
	}
}

func printString(s string, c chan string) {
	fmt.Println(s)
	c <- "Done printing."
}
*/

/*
package main

import "fmt"

func main() {
	c := make(chan string)

	for i := 0; i < 4; i++ {
		go printString("Hello there!", c)
	}

	for {
		fmt.Println(<-c)
	}
}

func printString(s string, c chan string) {
	fmt.Println(s)
	c <- "Done printing."
}
*/
