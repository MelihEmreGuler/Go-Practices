package main

func main() {
	dog := Dog{
		Name:  "Samson",
		Breed: "German Shephered",
	}
	PrintInfo(&dog)

	gorilla := Gorilla{
		Name:          "King Kong",
		Color:         "Black",
		NumberOfTeeth: 32,
	}
	PrintInfo(&gorilla)

	/*	eb := englishBot{}
		sb := spanishBot{}

		printGreeting(eb)
		printGreeting(eb)

		printGoodbye(eb)
		printGoodbye(eb)

		fmt.Println(sb)
	*/
}
