package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const portNumber string = ":8080"

// Home is the home page handler
func Home(responseWriter http.ResponseWriter, request *http.Request) {
	renderTemplate(responseWriter, "home.page.tmpl")
}

// About is the about page handler
func About(responseWriter http.ResponseWriter, request *http.Request) {
	renderTemplate(responseWriter, "about.page.tmpl")
}

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

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Starting application on port %s\n", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
