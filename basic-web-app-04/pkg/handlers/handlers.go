package handlers

import (
	"basic-web-app/pkg/render"
	"net/http"
)

// Home is the home page handler
func Home(responseWriter http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(responseWriter, "home.page.tmpl")
}

// About is the about page handler
func About(responseWriter http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(responseWriter, "about.page.tmpl")
}
