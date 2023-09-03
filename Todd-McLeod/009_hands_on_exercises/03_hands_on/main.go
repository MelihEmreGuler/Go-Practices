// 1. Take the previous program and change it so that:
// * func main uses http.Handle instead of http.HandleFunc

// Contstraint: Do not change anything outside of func main

// Hints:

// [http.HandlerFunc](https://godoc.org/net/http#HandlerFunc)
// ``` Go
// type HandlerFunc func(ResponseWriter, *Request)
// ```

// [http.HandleFunc](https://godoc.org/net/http#HandleFunc)
// ``` Go
// func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
// ```

// [source code for HandleFunc](https://golang.org/src/net/http/server.go#L2069)
// ``` Go
//
//	func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
//			mux.Handle(pattern, HandlerFunc(handler))
//	}
//
// ```
package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.Handle("/", http.HandlerFunc(foo))
	http.Handle("/dog/", http.HandlerFunc(bar))
	http.Handle("/me/", http.HandlerFunc(myName))
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	//io.WriteString(w, "foo ran")
	w.Write([]byte("foo ran\n")) // this is the same as io.WriteString

	w.WriteHeader(http.StatusOK) // this is the same as w.WriteHeader(200)

	x := req.URL.Path[1:] // this is the path that was requested
	charToRemove := "/"
	newString := strings.Replace(x, charToRemove, " ", -1) // this replaces the charToRemove with a space
	io.WriteString(w, newString)
}

func bar(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "bar ran")
}

func myName(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("something.gohtml")
	if err != nil {
		log.Fatalln("error parsing templete", err)
	}

	err = tpl.ExecuteTemplate(w, "something.gohtml", "Melih Emre Güler")
	if err != nil {
		log.Fatalln("error executing temlete", err)
	}
}
