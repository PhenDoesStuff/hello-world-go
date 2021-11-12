package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		n, err := fmt.Fprintf(w, "Hello, world!") // Fprintf returns a number of bytes written and error, so we have to do something with them (see below)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("Number of bytes written: %d", n)) // Sprintf allows for us to add different types into print statements using %d
	})

	_ = http.ListenAndServe(":8080", nil) // This is how we start a webserver in Go - Ignoring the error that comes back if there is one
}