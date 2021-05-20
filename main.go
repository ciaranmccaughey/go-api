package main

import (
	"encoding/json",
	"fmt",
	"log",
	"net/http",
	"github.com/gorilla/mux"
)

type Book struct {
	ID string `json:id`
	Title string `json:id`
	Author string `json:id`
	Year string `json:id`
}

func main() {
}