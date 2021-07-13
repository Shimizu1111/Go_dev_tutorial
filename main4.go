// // code:web3-13
// package main

// import (
// 	"html/template"
// 	"log"
// 	"net/http"
// 	"time"
// )

// // ①Embed htmlファイルに埋め込むデータ構造体
// type Emb struct {
// 	Title   string
// 	Message string
// 	Time    time.Time
// }

// func main() {
// 	port := "8080"

// 	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
// 	// ②
// 	http.HandleFunc("/", handleIndex)
// 	log.Printf("Server listening on http://localhost:%s/", port)
// 	log.Print(http.ListenAndServe(":"+port, nil))
// }

// func handleIndex(w http.ResponseWriter, r *http.Request) {
// 	// ③
// 	t, err := template.ParseFiles(
// 		"root/index.html",
// 		"root/template/header.html",
// 		"root/template/footer.html",
// 	)
// 	if err != nil {
// 		log.Fatalf("template error: %v", err)
// 	}
// 	// ④
// 	temp := Emb{"Hello Golang!", "こんにちは！", time.Now()}
// 	if err := t.Execute(w, temp); err != nil {
// 		log.Printf("failed to execute template: %v", err)
// 	}
// }

// // code:web3-13
// package main

// import (
// 	"html/template"
// 	"log"
// 	"net/http"
// 	"time"
// )

// // ①Embed htmlファイルに埋め込むデータ構造体
// type Emb struct {
// 	Title   string
// 	Message string
// 	Time    time.Time
// }

// func main() {
// 	port := "8080"

// 	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
// 	// ②
// 	http.HandleFunc("/", handleIndex)
// 	log.Printf("Server listening on http://localhost:%s/", port)
// 	log.Print(http.ListenAndServe(":"+port, nil))
// }

// func handleIndex(w http.ResponseWriter, r *http.Request) {
// 	// ③
// 	t, err := template.ParseFiles(
// 		"root/index.html",
// 		"root/template/header.html",
// 		"root/template/footer.html",
// 	)
// 	if err != nil {
// 		log.Fatalf("template error: %v", err)
// 	}
// 	// ④
// 	temp := Emb{"Hello Golang!", "こんにちは！", time.Now()}
// 	if err := t.Execute(w, temp); err != nil {
// 		log.Printf("failed to execute template: %v", err)
// 	}
// }

// code:web3-13
package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

// ①Embed htmlファイルに埋め込むデータ構造体
type Emb struct {
	Title   string
	Message string
	Time    time.Time
}

func main() {
	port := "8080"

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	// ②
	http.HandleFunc("/", handleIndex)
	log.Printf("Server listening on http://localhost:%s/", port)
	log.Print(http.ListenAndServe(":"+port, nil))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	// ③
	t, err := template.ParseFiles(
		"root/index.html",
		"root/template/header.html",
		"root/template/footer.html",
	)
	if err != nil {
		log.Fatalf("template error: %v", err)
	}
	// ④
	temp := Emb{"Hello Golang!", "こんにちは！", time.Now()}
	if err := t.Execute(w, temp); err != nil {
		log.Printf("failed to execute template: %v", err)
	}
}
