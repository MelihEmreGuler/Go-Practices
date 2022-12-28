package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	tpl, err := template.ParseFiles("one.gmao")

	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, nil)

	if err != nil {
		log.Fatalln(err)
	}

	tpl, err = tpl.ParseFiles("two.gmao", "wespa.gmao")
	if err != nil {
		log.Fatalln(err)
	}
}
