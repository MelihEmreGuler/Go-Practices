package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://amazon.com",
	}

	for _, link := range links {
		checkLink(link)
	}
}
func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		return //fonksiyonun icerisinde burdan sonra baska bir sey yapmamasi icin
	}
	fmt.Println(link, "its up!")
}
