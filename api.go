package main

import (
	"fmt"
	"log"
	"net/http"
)

func apiserver() {
	router := NewRouter()
	// Find a way to push what server a client is connecting to from the main api endpoint

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "gGoban API")
}
