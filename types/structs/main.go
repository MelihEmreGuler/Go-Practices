package main

import (
	"fmt"
	"time"
)

func main() {
	alex := Person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)

	var melih Person
	fmt.Println(melih)
	fmt.Printf("%+v\n", melih) //%+v : sahip oldugu tum tanimlamalarla birlikte bastir demek

	jim := Person{
		firstName: "Jim",
		lastName:  "Party",
		ContactInfo: ContactInfo{
			email:   "jim@gmail.com",
			zipCode: 18000,
		},
	}
	fmt.Println(jim)
	fmt.Printf("%+v\n", jim)

	jim.print()

	jim.updateName("Jimmy")

	jim.print()

	var melih2 = NewPerson{
		FirstName:   "Melih",
		LastName:    "Güler",
		PhoneNumber: "0535 000 00 00",
		Age:         25,
		BirthDate:   time.Now(),
	}

	fmt.Println(melih2)
}
