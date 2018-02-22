package main

import (
	"Cloud-Native-Go/api"
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", index)

	http.HandleFunc("/api/books", api.BookHandleFunc)
	http.HandleFunc("/api/books/", api.BookHandleFunc)

	http.ListenAndServe(port(), nil)
}

func port() string {
	port := os.Getenv("CNATIVE_PORT")
	if len(port) == 0 {
		port = "9090"
	}
	return ":" + port
}

func index(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	fmt.Fprintf(res, "Welcome to Cloud Native Go (Update).")
}
