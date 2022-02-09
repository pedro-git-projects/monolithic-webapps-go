package main

import (
	"net/http"
)

// Home is the home page handler
func Home(responseWriter http.ResponseWriter, request *http.Request) {
	renderTemplate(responseWriter, "home.page.tmpl")
}

// About is the about page handler
func About(responseWriter http.ResponseWriter, request *http.Request) {
	renderTemplate(responseWriter, "about.page.tmpl")
}
