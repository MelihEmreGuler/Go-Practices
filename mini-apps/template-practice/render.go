package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func renderTemplate(w http.ResponseWriter, tmpl string) {
	// ParseFiles returns a pointer to a template and an error
	parsedTemplate, _ := template.ParseFiles("template-practice/templates/" + tmpl)

	// Execute applies a parsed template to the specified data object, writing the output to writer w.
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template", err)
		return
	}

	fmt.Println(tmpl, "rendered successfully")
}
