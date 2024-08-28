package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/awalker125/go-api/handlers/person"
	"github.com/awalker125/go-api/middleware"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", middleware.Chain(Hello, middleware.Method("GET"), middleware.Logging()))

	handler := person.PersonHandler{Name: "bob"}

	chain1 := middleware.Chain(handler.ServeHTTP, middleware.Method("GET"), middleware.Logging())

	mux.Handle("/person", chain1)

	http.HandleFunc("/person", handler.ServeHTTP)
	http.HandleFunc("/person2", middleware.Chain(handler.ServeHTTP, middleware.Method("GET"), middleware.Logging()))
	log.Println("Starting server on 8080...")
	http.ListenAndServe(":8080", mux)
}
