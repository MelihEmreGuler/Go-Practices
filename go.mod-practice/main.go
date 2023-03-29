package main

import (
	"fmt"

	"github.com/gorilla/mux"
)

func main() {
	_ = mux.NewRouter()

	fmt.Println("Hello World")
}
