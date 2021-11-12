package main

import (
	"fmt"
	"net/http"

	"github.com/stephenmontague/hello-world-go/pkg/handlers"
)

const portNumber = ":8080"

// main is the main application function
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port %s", portNumber)

	_ = http.ListenAndServe(portNumber, nil) // This is how we start a webserver in Go - Ignoring the error that comes back if there is one
}