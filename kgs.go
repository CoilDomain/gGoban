package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/parnurzeal/gorequest"
)

var kgsapiurl = "https://www.gokgs.com/json/access"

type kgscredentials struct {
	Type     string `json:"type"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Locale   string `json:"locale"`
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJSON(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}

var request = gorequest.New()

func kgsLogin() int {

	// Define credentials
	login := &kgscredentials{
		Type:     "LOGIN",
		Name:     "ggoban",
		Password: "jason",
		Locale:   "en_US",
	}
	// Initiate connection
	resp, body, errs := request.Post(kgsapiurl).Send(login).End()
	if errs != nil {
		fmt.Println(errs)
		os.Exit(1)
	}
	fmt.Println("Login Status:", body)
	fmt.Println("Login Status:", resp.Status)
	fmt.Println("Login Headers:", resp.Header)

	return resp.StatusCode
}

func kgsPoll() (jsonbody string) {
	// Validate request succeeded
	if kgsLogin() == 200 {
		// Initiate GET for new data
		resp, body, errs := request.Get(kgsapiurl).End()
		if errs != nil {
			fmt.Println(errs)
			os.Exit(1)
		}

		jsonbody := body
		return jsonbody
		fmt.Println("Poll Status:", resp.Status)
		fmt.Println("Poll Headers:", resp.Header)
		fmt.Println("Poll Body:", body)
	} else {
		// Else failback and try reconnecting
		fmt.Println("Disconnected or Errored out")
		kgsLogin()
	}
	return
}
