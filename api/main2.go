package main

import (
	"fmt"
	"log"
	"net/http"
)

type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		handleIndex(w, r)
	case "/greeting":
		handleGreet(w, r)
	default:
		http.NotFound(w, r)
	}
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello1")
}

func handleGreet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello2")
}

func main() {
	port := "8080"
	mux := &MyMux{}
	log.Printf("http://localhost:%s/", port)
	log.Print(http.ListenAndServe(":"+port, mux))
}
