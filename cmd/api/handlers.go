package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/mjmichael73/go-books-api/internal/data"
)

func (app *application) handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}
	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	js = append(js, '\n')
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (app *application) handleGetCreateBooks(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		books := []data.Book{
			{
				ID:        1,
				CreatedAt: time.Now(),
				Title:     "The Darkening of Tristram",
				Published: 1998,
				Pages:     300,
				Genres:    []string{"Fiction", "Thriller"},
				Rating:    4.5,
				Version:   1,
			},
			{
				ID:        1,
				CreatedAt: time.Now(),
				Title:     "The Legacy of Deckard Cain",
				Published: 1998,
				Pages:     300,
				Genres:    []string{"Fiction", "Adventure"},
				Rating:    4.5,
				Version:   1,
			},
		}
		if err := app.writeJSON(w, http.StatusOK, books); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}
	if r.Method == http.MethodPost {
		fmt.Fprintln(w, "Added a new book to the reading list")
	}
}

func (app *application) handleGetUpdateBooks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		app.getBook(w, r)
	case http.MethodPut:
		app.updateBook(w, r)
	case http.MethodDelete:
		app.deleteBook(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (app *application) getBook(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/v1/books/"):]
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	book := data.Book{
		ID:        idInt,
		CreatedAt: time.Now(),
		Title:     "Echoes in the Darkness",
		Published: 2019,
		Pages:     300,
		Genres:    []string{"Fixtion", "Thriller"},
		Rating:    4.5,
		Version:   1,
	}
	if err := app.writeJSON(w, http.StatusOK, book); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (app *application) updateBook(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("v1/books/"):]
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Update the details of book with ID: %d", idInt)
}

func (app *application) deleteBook(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("v1/books/"):]
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Delete the details of book with ID: %d", idInt)
}
