package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})

	http.HandleFunc("/usuarios", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("Carregar usuarios"))
	})

	log.Fatal(http.ListenAndServe(":5001", nil))
}
