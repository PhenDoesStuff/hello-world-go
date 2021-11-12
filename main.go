package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(1, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page, and 1 + 2 is %d", sum))
}

// addValues adds two integers and returns the sum
func addValues(x, y int) int {
	return x + y
}

// main is the main application function
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

fmt.Println(fmt.Sprintf("Starting application onf port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil) // This is how we start a webserver in Go - Ignoring the error that comes back if there is one
}