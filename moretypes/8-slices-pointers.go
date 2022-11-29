package main

import "fmt"

func main() {
	names := [5]string{
		"melih",
		"emre",
		"guler",
		"mert",
		"senel",
	}
	fmt.Println(names)

	slicedNames1 := names[0:3]
	slicedNames2 := names[1:4]

	slicedNames1[2] = "tarik" //birinden degistirdigim tum hepsinden degismesine neden oldu

	fmt.Printf("names: %v slicedNames1: %v, slicedNames2: %v \n", names, slicedNames1, slicedNames2)

	slicedNames3 := names[3:] //ust limitin default olmasini istedgimiz yerde orayi bosluk birakabiliriz

	fmt.Println(slicedNames3)
}
