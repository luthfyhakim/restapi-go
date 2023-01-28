package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HandleHome).Methods("GET")

	port := 8111
	log.Println("Starting web on port", port)

	err := http.ListenAndServe(":8111", r)
	log.Fatal(err)
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	status := 200
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  status,
		"title":   "Home",
		"message": "Hello World!",
	})
}
