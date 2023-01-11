// ListenAndServe on port ":8080" using the default ServeMux.

// Use HandleFunc to add the following routes to the default ServeMux:

// "/"
// "/dog/"
// "/me/

// Add a func for each of the routes.

// Have the "/me/" route print out your name.

package main

import (
	"io"
	"net/http"
)

func foo(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "foo ran")
}
func bar(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "bar ran")
}
func me(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "My name is Melih")
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", bar)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":8080", nil)
}
