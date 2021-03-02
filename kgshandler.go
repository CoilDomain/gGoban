package main

import (
	"io"
	"net/http"
)

var kgsroutes = Route{
	"kgsapiindex",
	"GET",
	"/kgs",
	kgsapiindex,
}

func kgsapiindex(w http.ResponseWriter, r *http.Request) {
	jsonbody := kgsPoll()
	io.WriteString(w, jsonbody)
}
