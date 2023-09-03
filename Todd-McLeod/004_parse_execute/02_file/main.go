package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	tpl, err := template.ParseFiles("tpl.gohtml") // Parse the template

	if err != nil {
		log.Fatalln(err)
	}

	nf, err := os.Create("index.html") // Create a file

	if err != nil {
		log.Println("error creating file", err)
	}
	defer nf.Close() // Close the file

	err = tpl.Execute(nf, nil) // Execute the template to the file
 
	if err != nil {
		log.Fatalln(err)
	}
}
