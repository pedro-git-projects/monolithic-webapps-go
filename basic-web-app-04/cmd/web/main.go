package main

import (
	"basic-web-app/pkg/handlers"
	"fmt"
	"net/http"
)

const portNumber string = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port %s\n", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
