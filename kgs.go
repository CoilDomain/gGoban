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

type MyJsonName struct {
	Messages []struct {
		BlitzOk       bool   `json:"blitzOk"`
		ChannelID     int64  `json:"channelId"`
		Description   string `json:"description"`
		EstimatedRank string `json:"estimatedRank"`
		FastOk        bool   `json:"fastOk"`
		FreeOk        bool   `json:"freeOk"`
		Friends       []struct {
			FriendType string `json:"friendType"`
			User       struct {
				AuthLevel string `json:"authLevel"`
				Flags     string `json:"flags"`
				Name      string `json:"name"`
				Rank      string `json:"rank"`
			} `json:"user"`
		} `json:"friends"`
		Games []struct {
			Adjourned       bool   `json:"adjourned"`
			ChannelID       int64  `json:"channelId"`
			GameType        string `json:"gameType"`
			Global          bool   `json:"global"`
			Handicap        int64  `json:"handicap"`
			InitialProposal struct {
				GameType string `json:"gameType"`
				Nigiri   bool   `json:"nigiri"`
				Players  []struct {
					Role string `json:"role"`
					User struct {
						Flags string `json:"flags"`
						Name  string `json:"name"`
						Rank  string `json:"rank"`
					} `json:"user"`
				} `json:"players"`
				Rules struct {
					ByoYomiPeriods int64   `json:"byoYomiPeriods"`
					ByoYomiStones  int64   `json:"byoYomiStones"`
					ByoYomiTime    int64   `json:"byoYomiTime"`
					Komi           float64 `json:"komi"`
					MainTime       int64   `json:"mainTime"`
					Rules          string  `json:"rules"`
					Size           int64   `json:"size"`
					TimeSystem     string  `json:"timeSystem"`
				} `json:"rules"`
			} `json:"initialProposal"`
			Komi      float64 `json:"komi"`
			MoveNum   int64   `json:"moveNum"`
			Name      string  `json:"name"`
			Observers int64   `json:"observers"`
			Over      bool    `json:"over"`
			Players   struct {
				Black struct {
					Flags string `json:"flags"`
					Name  string `json:"name"`
					Rank  string `json:"rank"`
				} `json:"black"`
				ChallengeCreator struct {
					Flags string `json:"flags"`
					Name  string `json:"name"`
					Rank  string `json:"rank"`
				} `json:"challengeCreator"`
				Owner struct {
					AuthLevel string `json:"authLevel"`
					Flags     string `json:"flags"`
					Name      string `json:"name"`
					Rank      string `json:"rank"`
				} `json:"owner"`
				White struct {
					AuthLevel string `json:"authLevel"`
					Flags     string `json:"flags"`
					Name      string `json:"name"`
					Rank      string `json:"rank"`
				} `json:"white"`
			} `json:"players"`
			Private  bool    `json:"private"`
			RoomID   int64   `json:"roomId"`
			Saved    bool    `json:"saved"`
			Score    float64 `json:"score"`
			Size     int64   `json:"size"`
			Uploaded bool    `json:"uploaded"`
		} `json:"games"`
		HumanOk         bool   `json:"humanOk"`
		JSONClientBuild string `json:"jsonClientBuild"`
		MaxHandicap     int64  `json:"maxHandicap"`
		MediumOk        bool   `json:"mediumOk"`
		Owners          []struct {
			AuthLevel string `json:"authLevel"`
			Flags     string `json:"flags"`
			Name      string `json:"name"`
			Rank      string `json:"rank"`
		} `json:"owners"`
		Playbacks []struct {
			DateStamp   string `json:"dateStamp"`
			GameSummary struct {
				GameType string  `json:"gameType"`
				Handicap int64   `json:"handicap"`
				Komi     float64 `json:"komi"`
				Players  struct {
					Owner struct {
						Flags string `json:"flags"`
						Name  string `json:"name"`
						Rank  string `json:"rank"`
					} `json:"owner"`
				} `json:"players"`
				Revision  int64  `json:"revision"`
				Score     string `json:"score"`
				Size      int64  `json:"size"`
				Timestamp string `json:"timestamp"`
			} `json:"gameSummary"`
			SubscribersOnly bool `json:"subscribersOnly"`
		} `json:"playbacks"`
		RankedOk               bool `json:"rankedOk"`
		RobotOk                bool `json:"robotOk"`
		RoomCategoryChannelIds struct {
			Clubs      int64 `json:"CLUBS"`
			Friendly   int64 `json:"FRIENDLY"`
			Lessons    int64 `json:"LESSONS"`
			Main       int64 `json:"MAIN"`
			National   int64 `json:"NATIONAL"`
			Temporary  int64 `json:"TEMPORARY"`
			Tournament int64 `json:"TOURNAMENT"`
		} `json:"roomCategoryChannelIds"`
		Rooms []struct {
			Category  string `json:"category"`
			ChannelID int64  `json:"channelId"`
			Name      string `json:"name"`
		} `json:"rooms"`
		Subscriptions []struct {
			End   string `json:"end"`
			Start string `json:"start"`
		} `json:"subscriptions"`
		Type       string `json:"type"`
		UnrankedOk bool   `json:"unrankedOk"`
		Users      []struct {
			AuthLevel string `json:"authLevel"`
			Flags     string `json:"flags"`
			Name      string `json:"name"`
			Rank      string `json:"rank"`
		} `json:"users"`
		VersionBugfix int64 `json:"versionBugfix"`
		VersionMajor  int64 `json:"versionMajor"`
		VersionMinor  int64 `json:"versionMinor"`
		You           struct {
			Flags string `json:"flags"`
			Name  string `json:"name"`
		} `json:"you"`
	} `json:"messages"`
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

func kgsPoll() {
	// Validate request succeeded
	if kgsLogin() == 200 {
		// Initiate GET for new data
		resp, body, errs := request.Get(kgsapiurl).End()
		if errs != nil {
			fmt.Println(errs)
			os.Exit(1)
		}
		fmt.Println("Poll Status:", resp.Status)
		fmt.Println("Poll Headers:", resp.Header)
		fmt.Println("Poll Body:", body)
	} else {
		// Else failback and try reconnecting
		fmt.Println("Disconnected or Errored out")
		kgsLogin()
	}
}
