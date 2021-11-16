package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/stephenmontague/hello-world-go/pkg/config"
	"github.com/stephenmontague/hello-world-go/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/About", http.HandlerFunc(handlers.Repo.About))

	return mux
}