// passing data with QueryString
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/search", search)
	http.ListenAndServe(":8080", nil)
}

func search(w http.ResponseWriter, r *http.Request) {
	q := r.FormValue("q")
	bih := r.FormValue("bih")
	biw := r.FormValue("biw")
	hl := r.FormValue("hl")
	sxsrf := r.FormValue("sxsrf")
	source := r.FormValue("source")
	ei := r.FormValue("ei")

	data := "q=" + q + "&bih=" + bih + "&biw=" + biw + "&hl=" + hl + "&sxsrf=" + sxsrf + "&source=" + source + "&ei=" + ei

	fmt.Fprintln(w, data)
}

//https://www.google.com/search?q=melih+emre+g%C3%BCler&bih=833&biw=1467&hl=tr&sxsrf=AJOqlzVwPE-4uT-LZL9XgcxN1z7gDsk16Q%3A1673718100020&source=hp&ei=U-nCY6TqO_He1sQP5faZqAU
