package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		bytesWritten, err := fmt.Fprintf(writer, "Hello World")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("The numberof bytes written was %d\n", bytesWritten)
	})

	/*
	 * func ListenAndServe(addr string, handler Handler) error
	 * ListenAndServe listens on the TCP network address addr
	 * and then calls Serve with handler to handle requests on incoming connections.
	 * The handler is typically nil, in which case the DefaultServeMux is used.
	 */
	_ = http.ListenAndServe(":8080", nil)
}
