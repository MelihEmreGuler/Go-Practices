package main

import "fmt"

const numPool = 10

func CalculateValue(intChan chan int) {
	randomNumber := RandomNumber(numPool)
	intChan <- randomNumber // send random number to channel, it provides a way for one goroutine to pass data to another goroutine
}

func main() {
	intChan := make(chan int) // create a channel that can hold an integer
	defer close(intChan)      // close the channel when the main function returns to avoid a deadlock(https://golang.org/doc/effective_go.html#channels)

	go CalculateValue(intChan) // start a goroutine to calculate a random number and send it to the channel

	randNum := <-intChan // receive the random number from the channel and assign it to a variable

	fmt.Printf("new random number is: %v", randNum)
}

/*
So that's just a really useful way of passing information from one package to another package
or one part of your program to another part of your program.
*/
