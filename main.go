package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// main is the main application function
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil) // This is how we start a webserver in Go - Ignoring the error that comes back if there is one
}