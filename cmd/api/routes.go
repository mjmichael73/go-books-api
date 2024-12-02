package main

import "net/http"

func (app *application) route() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", app.handleHealthCheck)
	mux.HandleFunc("/v1/books", app.handleGetCreateBooks)
	mux.HandleFunc("/v1/books/", app.handleGetUpdateBooks)
	return mux
}
