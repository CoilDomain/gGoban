package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func apiserver() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/kgs", kgsapiindex)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "gGoban API")
}

func kgsapiindex(w http.ResponseWriter, r *http.Request) {
	jsonbody := kgsPoll()
	io.WriteString(w, jsonbody)
}
