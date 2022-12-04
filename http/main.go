package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	//fmt.Println(resp)
	/*
		bs := make([]byte, 99999)

		resp.Body.Read(bs) //biz read fonksiyonun icerisine ihtiyaci olan buyuklukten daha buyuk bir bir slice verdik o bu slice'i bolecek

		fmt.Println(string(bs)) //byte slice' ler aslında bir string gibidir.
		//html dosyasini elde ettik

		ioutil.WriteFile("html", bs, 0666)
	*/

	io.Copy(os.Stdout, resp.Body) //bir ustte yorum bloguna aldigim satirlarla tamamen aynı sekilde google nin html dosyasini konsola bastiriyoruz.
	// go dilinin kisayolu

	lw := logWriter{}

	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrotre this many bytes: ", len(bs))
	return len(bs), nil
}
