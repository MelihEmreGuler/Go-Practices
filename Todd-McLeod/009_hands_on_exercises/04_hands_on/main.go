package main

import (
	//"io"
	"fmt"
	"log"
	"net/http"
)

type messageHandler struct {
	message string
}

func (m *messageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//io.WriteString(w, m.message)
	// w.Write([]byte(m.message))
	// fmt.Fprintf(w, m.message)
	fmt.Fprintln(w, m.message)
}

func main() {
	mux := http.NewServeMux()

	m1 := &messageHandler{"first message"}
	m2 := &messageHandler{"second message"}

	mux.Handle("/first/", m1)
	mux.Handle("/second/", m2)

	log.Println("Listening...")

	http.ListenAndServe(":8080", mux)

}
