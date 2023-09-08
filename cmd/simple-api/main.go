package main

import (
	"io"
	"log"
	"net/http"
)

func getHello(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.WriteHeader(http.StatusOK)
		log.Printf("/hello %v %v", r.Method, http.StatusOK)

		nameValue := r.URL.Query().Get("name")
		if nameValue != "" {
			if _, err := io.WriteString(w, "Hello "+nameValue+"!"); err != nil {
				log.Fatal(err)
			}
		} else {
			if _, err := io.WriteString(w, "Hello!"); err != nil {
				log.Fatal(err)
			}
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("/hello %v %v", r.Method, http.StatusBadRequest)
		if _, err := io.WriteString(w, "Sorry we accept only GET requests"); err != nil {
			log.Fatal(err)
		}
	}
}

// func getHelloQuery(w http.ResponseWriter, r *http.Request) {
// 	nameValue := r.URL.Query().Get("name")
// 	if nameValue != "" {
// 		if _, err := io.WriteString(w, "Hello "+nameValue+"!"); err != nil {
// 			log.Fatal(err)
// 		}
// 	}
// }

func main() {
	http.HandleFunc("/hello", getHello)
	// http.HandleFunc("/hello?name", getHelloQuery)
	// http.HandleFunc("/")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
