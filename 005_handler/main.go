package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) { // any type that has a ServeHTTP method is a Handler
	fmt.Fprintln(w, "Any code you want in this func")
}

func main() {
	var d hotdog // d is a Handler
	http.ListenAndServe(":8080", d) //postman localhost:8080 will print "Any code you want in this func"

}