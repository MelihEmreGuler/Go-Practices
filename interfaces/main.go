package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func (englishBot) getGreeting() string { //func (eb englishBot) getGreeting() string {} yazmamiz arasinda bir fark yok.
	//eger kullanmicaksak eb kopyasini siliebiliriz.
	return "Hi There!"
}
func (spanishBot) getGreeting() string {
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

/*
func printGreeting(eb englishBot) { //go da method overloading olmadigi icin aynı isime sahip fonksiyonlar hata vericektir
	fmt.Println(eb.getGreeting())
}
func printGreeting(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}

*/
 