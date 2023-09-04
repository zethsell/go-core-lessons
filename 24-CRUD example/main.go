package main

import (
	"crud/data/property"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/properties", property.InsertProperty).Methods(http.MethodPost)
	router.HandleFunc("/properties", property.ListProperties).Methods(http.MethodGet)
	router.HandleFunc("/properties/{id}", property.ShowProperty).Methods(http.MethodGet)
	router.HandleFunc("/properties/{id}", property.UpdateProperty).Methods(http.MethodPut)
	router.HandleFunc("/properties/{id}", property.DeleteProperty).Methods(http.MethodDelete)

	fmt.Println("server listening on port 6000")
	log.Fatal(http.ListenAndServe(":6000", router))
}
