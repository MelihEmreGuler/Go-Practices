package main

import "fmt"

type bot interface {
	getGreeting() string
	getGoodbye() string
}

type englishBot struct{}
type spanishBot struct{}

func (englishBot) getGreeting() string { //func (eb englishBot) getGreeting() string {} yazmamiz arasinda bir fark yok.
	//eger kullanmicaksak eb kopyasini siliebiliriz.
	return "Hi There!"
}
func (spanishBot) getGreeting() string {
	return "Hola!"
}
func (englishBot) getGoodbye() string {
	return "Goodbye!"
}

/*
func (spanishBot) getGoodbye() string { //it is necessary to implement all methods of interface
	return "Adios!"
}
*/

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
func printGoodbye(b bot) {
	fmt.Println(b.getGoodbye())
}

/*
func printGreeting(eb englishBot) { //go da method overloading olmadigi icin aynı isime sahip fonksiyonlar hata vericektir
	fmt.Println(eb.getGreeting())
}
func printGreeting(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}

*/
