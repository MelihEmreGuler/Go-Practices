package main

import "fmt"

type person struct {
	firstName   string
	lastName    string
	contactInfo //it is same with contactInfo: contactInfo
}
type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)

	var melih person
	fmt.Println(melih)
	fmt.Printf("%+v\n", melih) //%+v : sahip oldugu tum tanimlamalarla birlikte bastir demek

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 18000,
		},
	}
	fmt.Println(jim)
	fmt.Printf("%+v\n", jim)

	jim.print()

	jim.updateName("Jimmy")

	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateName(newFirstName string) { //eger fonksiyonun tuttugu referansi bir pointer olarak isaretlersek bu referansata yaptigimiz degisiklikler
	//fonksiyonu cagirdigimiz personun uzerinde gerceklesecek. Pointer secmezsek personun firstName' i degismeyecek. 
	p.firstName = newFirstName
}
