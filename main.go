package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/test", handler).Methods("PUT").Methods("PATCH")

	server := &http.Server{
		Addr: ":9710",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}