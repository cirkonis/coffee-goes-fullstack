package main

import (
	"coffee-goes-fullstack/pkg/routers"
	"encoding/json"
	"log"
	"net/http"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBookByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id")
	for _, book := range books {
		if book.ID == id {
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	http.NotFound(w, r)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id")
	var book Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for i, b := range books {
		if b.ID == id {
			books[i] = book
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	http.NotFound(w, r)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id")
	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	http.NotFound(w, r)
}

func main() {
	//delete this shit
	books = append(books, Book{ID: "1", Title: "The Catcher in the Rye", Author: "J.D. Salinger"})
	books = append(books, Book{ID: "2", Title: "To Kill a Mockingbird", Author: "Harper Lee"})

	http.HandleFunc("/books/create", createBook)
	http.HandleFunc("/books/update", updateBook)
	http.HandleFunc("/books/delete", deleteBook)

	// Start server
	r := routers.Init()
	log.Fatal(http.ListenAndServe(":8080", r))
}
