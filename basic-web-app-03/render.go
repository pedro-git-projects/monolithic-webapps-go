package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func renderTemplate(responseWriter http.ResponseWriter, tmpl string) {
	// ParseFiles returns pointer to a template and an error
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	// takes a responseWriter and data
	// Execute applies a parsed template to the specified data object
	// writing the output to wr. It only returns an error
	err := parsedTemplate.Execute(responseWriter, nil)
	if err != nil {
		fmt.Println("Error parsing template: ", err)
		return
	}
}
