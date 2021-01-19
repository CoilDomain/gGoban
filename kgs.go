package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

var kgsapiurl = "https://www.gokgs.com/json/access"

type kgscredentials struct {
	Type     string `json:"type"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Locale   string `json:"locale"`
}

type messages struct {
	VersionBugfix   string `json:"versionBugfix"`
	Type            string `json:"type"`
	VersionMajor    string `json:"versionMajor"`
	VersionMinor    string `json:"versionMinor"`
	JsonClientBuild string `json:"jsonClientBuild"`

	subscriptions struct {
		Start string `json:"start"`
		End   string `json:"end"`
	}

	rooms struct {
		Category   string `json:"category"`
		ChannnelID int32  `json:"channelID"`
	}

	roomCategoryChannelIds struct {
		NATIONAL   int32 `json:"NATIONAL"`
		TOURNAMENT int32 `json:"TOURNAMENT"`
		FRIENDLY   int32 `json:"FRIENDLY"`
		TEMPORARY  int32 `json:"TEMPORARY"`
		MAIN       int32 `json:"MAIN"`
		LESSONS    int32 `json:"LESSONS"`
		CLUBS      int32 `json:"CLUBS"`
	}

	friends struct {
		FriendType string `json:"friendType"`
		User       struct {
			Name  string `json:"name"`
			Flags string `json:"flags"`
			Rank  string `json:"rank"`
		}
	}
	Games struct {
		Adjourned string `json:"adjourned"`
		GameType  string `json:"gameType"`
		Komi      int64  `json:"komi"`
		Size      int32  `json:"size"`
		Handicap  int32  `json:"handicap"`
		Players   struct {
			Owner struct {
				Name  string `json:"name"`
				Flags string `json:"flags"`
				Rank  string `json:"rank"`
			}
			White struct {
				Name  string `json:"name"`
				Flags string `json:"flags"`
				Rank  string `json:"rank"`
			}
			Black struct {
				Name  string `json:"name"`
				Flags string `json:"flags"`
				Rank  string `json:"rank"`
			}
		}
		MoveNum    int32  `json:"moveNum"`
		Global     string `json:"global"`
		ChannnelID int32  `json:"channnelID"`
		Observers  int32  `json:"observers"`
		RoomID     int32  `json:"roomID"`
		Over       string `json:"over"`
		Score      string `json:"score"`
		Saved      string `json:"saved"`
	}
}

func kgsLogin() {

	// Define credentials
	login := kgscredentials{
		Type:     "LOGIN",
		Name:     "lazerspine",
		Password: "",
		Locale:   "en_US",
	}

	// Initialize http client
	client := &http.Client{}

	// login to json
	jsondata, err := json.Marshal(login)
	if err != nil {
		panic(err)
	}

	// set method, url and body
	req, err := http.NewRequest(http.MethodPost, kgsapiurl, bytes.NewBuffer(jsondata))
	if err != nil {
		panic(err)
	}

	// Set request header content type for json
	req.Header.Set("content-type", "application/json;charset=UTF-8")
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
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
