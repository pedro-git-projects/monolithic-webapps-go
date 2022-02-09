package main

import (
	"fmt"
	"net/http"
)

const portNumber string = ":8080"

// to handle requests, a function must take a response writer and a pointer to a request

// Home is the home page handler
func Home(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(responseWriter, "The home page")
}

// About is the about page handler
func About(responseWriter http.ResponseWriter, request *http.Request) {
	sum := addValues(2, 2)
	fmt.Fprintf(responseWriter, "The about page and 2 + 2 = %d", sum)
}

// addValues adds two integers and returns the sum
func addValues(x, y int) int {
	return x + y
}
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Starting application on port %s\n", portNumber)
	/*
	 * func ListenAndServe(addr string, handler Handler) error
	 * ListenAndServe listens on the TCP network address addr
	 * and then calls Serve with handler to handle requests on incoming connections.
	 * The handler is typically nil, in which case the DefaultServeMux is used.
	 */
	_ = http.ListenAndServe(portNumber, nil)
}
