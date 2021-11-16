package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/stephenmontague/hello-world-go/pkg/config"
	"github.com/stephenmontague/hello-world-go/pkg/handlers"
	"github.com/stephenmontague/hello-world-go/pkg/render"
)

const portNumber = ":8081"

// main is the main application function
func main() {
	var app config.AppConfig
	
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Starting application on port %s", portNumber)

	_ = http.ListenAndServe(portNumber, nil) // This is how we start a webserver in Go - Ignoring the error that comes back if there is one
}