package main

import (
	"encoding/json",
	"fmt",
	"log",
	"net/http",
	"github.com/gorilla/mux",
	"github.com/subosito/gotenv"
)

type Book struct {
	ID int `json:id`
	Title string `json:id`
	Author string `json:id`
	Year string `json:id`
}

var books []Book;

func main() {
	router := mux.NewRouter()
    router.HandleFunc("/books", getBooks).Methods.("GET");
    router.HandleFunc("/books/{id}", getBook).Methods.("GET");
    router.HandleFunc("/books", addBook).Methods.("POST");
    router.HandleFunc("/books", updateBook).Methods.("PUT");
    router.HandleFunc("/books/{id}", removeBook).Methods.("DELETE");
}