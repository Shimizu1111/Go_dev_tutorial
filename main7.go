package main

import (
	"log"
	"net/http"
	"text/template"
	"time"
)

type Emb struct {
	Title   string
	Message string
	Time    time.Time
}

func main() {
	port := "8080"

	http.HandleFunc("/", handleIndex)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	log.Printf("http://localhost:%s", port)
	log.Print(http.ListenAndServe(":"+port, nil))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(
		"root/index.html",
		"root/template/header.html",
		"root/template/footer.html",
	)

	if err != nil {
		log.Fatalf("template error: %v", err)
	}

	temp := Emb{
		Title:   "Hello Golang",
		Message: "こんにちは",
		Time:    time.Now(),
	}

	if err := t.Execute(w, temp); err != nil {
		log.Printf("failed to  execute template: %v", err)
	}
}
