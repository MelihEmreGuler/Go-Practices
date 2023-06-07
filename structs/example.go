package main

import (
	"fmt"
	"time"
)

type Person struct {
	firstName   string
	lastName    string
	ContactInfo //it is same with contactInfo: contactInfo
}
type ContactInfo struct {
	email   string
	zipCode int
}

type NewPerson struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func (p Person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *Person) updateName(newFirstName string) { //eger fonksiyonun tuttugu referansi bir pointer olarak isaretlersek bu referansata yaptigimiz degisiklikler
	//fonksiyonu cagirdigimiz personun uzerinde gerceklesecek. Pointer secmezsek personun firstName' i degismeyecek.
	p.firstName = newFirstName
}
