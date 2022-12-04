package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) print() { // like a method, we can use .print() on every variable which its type is deck (receiver function)
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ", ") //aralarına virgul koyarak string slice' ini tek bir stringe donusturuyor.
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFormFile(fileName string) deck {
	//bs yi tek bir string gibi dusunebiliriz.
	//byte string and error
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return deck(strings.Split(string(bs), ", ")) //sirasiyla byte'i stringe, stringi string dizesine string dizesini deck e cevirdim
}

func (d deck) shuffle() {

	r := rand.New(rand.NewSource(time.Now().UnixNano())) //tam anlamıyla rasgele bir random degiskeni uretme yontemimiz
	
	for index := range d {
		
		newPosition := r.Intn(len(d))                        //random degiskenini int deger icin kullanma yontemimiz

		d[index], d[newPosition] = d[newPosition], d[index]
	}
}
