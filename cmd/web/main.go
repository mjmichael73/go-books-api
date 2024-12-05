package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/mjmichael73/go-books-api/internal/models"
)

type application struct {
	readinglist *models.ReadingListModel
}

func main() {
	addr := flag.String("addr", ":8001", "Http Network Server")
	endpoint := flag.String("endpoint", "http://localhost:8000/v1/books", "Endpoint for the readinglist web service")
	app := &application{
		readinglist: &models.ReadingListModel{Endpoint: *endpoint},
	}
	srv := &http.Server{
		Addr:    *addr,
		Handler: app.routes(),
	}
	log.Printf("Starting the server on %s", *addr)
	err := srv.ListenAndServe()
	log.Fatal(err)
}
