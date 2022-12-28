package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://golang.org",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://amazon.com",
	}

	c := make(chan string) //channel created

	for _, link := range links {
		go checkLink(link, c) //gorutine created
	}

	//fmt.Println(<-c) //channeldan bir deger aliyoruz
	//bir mesaj alinmasini istemeye (fmt.Println(<-c) gibi) blocking adi verilir.
	//channel degiskeni bir deger kazandiktan sonra programa beni kullanmak isteyen bir yer var mi diye sorar.
	//eger varse oraya ilerler ve orada kullanilir.
	//program oradan asagiya dogru devam eder ve yukarida atlanan islemler bloke olur.
	//eger diger rutinlerde atanacak degerleri yakalayacak bir islem koyarsak program calismaya devam eder.
	//fmt.Println(<-c) //channeldan bir deger daha aliyoruz
	//fmt.Println(<-c) //channeldan bir deger daha aliyoruz
	//fmt.Println(<-c) //channeldan bir deger daha aliyoruz

	/*
		for i := 0; i < len(links); i++ { //yukaridaki satirlari bu sekilde de yapabiliriz
			fmt.Println(<-c) //channeldan bir deger aliyoruz
		}
	*/
	/*
		//surekli olarak bu sitelerin durumunu kontrol etmek istiyoruz
		for {
			go checkLink(<-c, c)
		}
	*/
	//bu sekilde de yapabiliriz
	for l := range c { //channeldan deger alinirken range kullanilabilir
		go func(link string) { //anonim fonksiyon //go literal
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l) //anonim fonksiyonu cagirirken l degerini kullanmak icin
	}
}
func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link //channela bir deger atiyoruz
		return    //fonksiyonun icerisinde burdan sonra baska bir sey yapmamasi icin
	}
	fmt.Println(link, "its up!")
	c <- link //channela bir deger atiyoruz
}

//channel <- 5 //Send the value 5 into this channel
//myNumber <- channel //Receive the value from the channel and assign it to myNumber
//fmt.Println(<-channel) //Receive the value from the channel and print it out
