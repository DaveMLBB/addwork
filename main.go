package main

import (
	"fmt"
	"log"
	"net/http"
)

// Handler che aggiunge le intestazioni CORS
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		fmt.Fprintf(w, "Hello, World!")
	})

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
