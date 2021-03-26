package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"phrases-server/handlers"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/reviewPhrases", handlers.GetReviewPhrasesEndpoint).Methods("GET")
	fmt.Println("Starting server on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", router))
}