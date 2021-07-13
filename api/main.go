package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world2")
}

func handleGreet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello guest")
}

func main() {
	port := "8080"
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/greeting", handleGreet)
	log.Printf("http://localhost:%s", port)
	//log.Printf("http://localhost:%s/", port)
	//nil の場合はDefaultServeMuxが適用される
	http.ListenAndServe(":"+port, nil)
}

// package main

// // ①
// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// // ②
// func handleIndex(w http.ResponseWriter, r *http.Request) {
//     fmt.Fprintf(w, "welcome golang server.")
// }

// func handleGreet(w http.ResponseWriter, r *http.Request) {
//     fmt.Fprintf(w, "hello guest!")
// }
// // ③
// func main() {
// 	port := "8080"
// 	http.HandleFunc("/", handleIndex)
// 	http.HandleFunc("/greeting", handleGreet)
// 	log.Printf("Server listening on http://localhost:%s/", port)
// 	// nil -> DefaultServeMuxが適用される
// 	log.Print(http.ListenAndServe(":"+port, nil))
// }
