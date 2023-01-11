// 1. Take the previous program in the previous folder and change it so that:
// * a template is parsed and served
// * you pass data into the template

package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", bar)
	http.HandleFunc("/me/", myName)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "foo ran")
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
