package main

import "fmt"

func main() {
	/*
		cards1 := deck{"asd", "dfg", "ghj"}
		cards1 = append(cards1, "jkl")
		cards1.print()
	*/

	/*
		cards2 := newDeck()
		cards2.print()

		hand, remainingCards := deal(cards2, 5)

		hand.print()
		remainingCards.print()
	*/
	/*
		greeting := "Hi there!"

		fmt.Println([]byte(greeting))
	*/
	/*
		cards3 := newDeck()

		fmt.Println(cards3.toString())

		fmt.Println("***************")

		cards3.saveToFile("my_cards")

		cards4 := newDeckFormFile("my_cards")
	*/
	cards5 := newDeck()

	cards5.print()

	cards5.shuffle()
	fmt.Println("Shuffled")

	cards5.print()
}
