package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

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
	handler.Name = "bill"

	chain1 := middleware.Chain(handler.ServeHTTP, middleware.Method("GET"), middleware.Logging())
	chain2 := middleware.Chain(person.Hello2, middleware.Method("GET"), middleware.Logging())
	chain3 := middleware.Chain(person.TimeHandler(time.RFC1123).ServeHTTP, middleware.Method("GET"), middleware.Logging())
	chain4 := middleware.Chain(person.TimeHandler2(time.RFC1123), middleware.Method("GET"), middleware.Logging())

	mux.Handle("/person", chain1)
	mux.Handle("/person2", chain2)
	mux.Handle("/time", chain3)
	mux.Handle("/time2", chain4)

	log.Println("Starting server on 8080...")
	http.ListenAndServe(":8080", mux)
}
