package main

import (
	"errors"
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

// Divide is the divide page handler, it writes a float32 that is the quotient of the divideValues function if
// the divisor is greater or different than 0 and it writes an eror message otherwise
func Divide(responseWriter http.ResponseWriter, request *http.Request) {
	quotient, err := divideValues(100.0, 10.0)
	if err != nil {
		fmt.Fprintf(responseWriter, "%v", err)
	} else {
		fmt.Fprintf(responseWriter, "The quotient of the division is %.2f", quotient)
	}
}

func divideValues(dividend, divisor float32) (float32, error) {
	if divisor <= 0 {
		err := errors.New("Cannot divide by zero!")
		return 0, err
	}
	quotient := dividend / divisor
	return quotient, nil

}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Printf("Starting application on port %s\n", portNumber)
	/*
	 * func ListenAndServe(addr string, handler Handler) error
	 * ListenAndServe listens on the TCP network address addr
	 * and then calls Serve with handler to handle requests on incoming connections.
	 * The handler is typically nil, in which case the DefaultServeMux is used.
	 */
	_ = http.ListenAndServe(portNumber, nil)
}
