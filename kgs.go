package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var kgsapiurl = "https://www.gokgs.com/json-cors/access"

type kgscredentials struct {
	Type     string `json:"type"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Locale   string `json:"locale"`
}

func kgsLogin() {

	// Define credentials
	login := kgscredentials{
		Type:     "LOGIN",
		Name:     "ggoban",
		Password: "jason",
		Locale:   "en_US",
	}

	// Initialize http client
	client := &http.Client{}

	// login to json
	json, err := json.Marshal(login)
	if err != nil {
		panic(err)
	}

	// set method, url and body
	req, err := http.NewRequest(http.MethodPost, kgsapiurl, bytes.NewBuffer(json))
	if err != nil {
		panic(err)
	}

	// Set request header content type for json
	req.Header.Set("Context-Type", "application/json; charset-utf-8")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}

func kgsPoll() {

}

func kgsSendMesg() {

}

func kgsRecMesg() {

}

func kgsNewGameReq() {

}

func kgsListGames() {

}
