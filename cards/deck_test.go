package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d)) //kullanacagimiz len(d) benzeri ifadeleri %v gibi place holders
		//kullanarak stringin icerisine dahil etmemiz gerek yoksa "t.Errorf("Expected deck length of 16, but got", len(d))" hatasi aliriz.
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected deck first card of Ace of Spades, but got %v", d[0])
	}
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected deck last card of Four of Clubs, but got %v", d[len(d)-1])
	}
}
func TestSaveToFileAndNewDeckFormFile(t *testing.T) {
	os.Remove("_decktesting") // alt cizgi ile baslatmamizin sebebi bu dosyanin gecici bir dosya oldugunun anlasilmasi
	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFormFile("_decktesting")
	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(loadedDeck)) //bu noktada testten basarisiz olarak ayrilsak bile
		//test fonksiyonun sonuna kadar gerceklestirecek ve gecici dosyayi silecek.
	}
	os.Remove("_decktesting")
}
