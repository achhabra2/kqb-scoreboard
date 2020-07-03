package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"fyne.io/fyne/app"
)

const (
	IglAPIURL        = "https://indy-gaming-league-api.herokuapp.com/api/circuits/"
	IglWestID        = "5e4c6b5178d46abdfeb49e71"
	IglEastID        = "5e4b295560f132acbb31b8f5"
	PlaceholderImage = "https://bulma.io/images/placeholders/128x128.png"
	KQBAvatarImage   = "/avatar.png"
)

type Team struct {
	Name  string
	Img   string
	Div   int
	Tier  int
	Stats Stats
}

type Stats struct {
	GamesWon    int
	GamesLost   int
	MatchesWon  int
	MatchesLost int
}

var h Team = Team{"Blue Team", PlaceholderImage, 1, 1, Stats{1, 1, 1, 1}}
var a Team = Team{"Gold Team", PlaceholderImage, 1, 1, Stats{1, 1, 1, 1}}
var s Scoreboard = Scoreboard{&h, &a, 0, 0, 0, 0}

func setupLogs() {
	f, err := os.OpenFile("./output.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}

	// don't forget to close it
	// defer f.Close()

	// assign it to the standard logger
	log.SetOutput(f)
}

func main() {
	// Configure log output to file
	setupLogs()

	// Setup HTTP handler and websocket handler
	StartHTTPServer()

	// Create Fyne App
	myApp := app.New()
	myApp.SetIcon(resourceIconPng)
	_ = GameType(myApp)
	myApp.Run()

	// Called before application exist
	tidyUp()
}

func GetTeamInfo(url string, c chan []Team) {
	log.Println("Fetching Team Information from IGL...")
	resp, err := http.Get(url)
	if err != nil {
		// handle error
		log.Println("Could not fetch team info from IGL")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var result IGLApiResponse
	json.Unmarshal([]byte(body), &result)
	teams := []Team{}
	for _, v := range result.Data {
		team := Team{}
		t := v.Team
		s := v.Stats

		team.Name = t.FormattedName
		team.Tier = int(t.Tier)
		team.Div = int(t.Div)
		if team.Img != "" {
			team.Img = t.Logo
		} else {
			team.Img = KQBAvatarImage
		}
		team.Stats.GamesWon, _ = strconv.Atoi(s.GamesWon)
		team.Stats.GamesLost, _ = strconv.Atoi(s.GamesLost)
		team.Stats.MatchesWon, _ = strconv.Atoi(s.MatchesWon)
		team.Stats.MatchesLost, _ = strconv.Atoi(s.MatchesLost)
		teams = append(teams, team)
	}

	c <- teams
}

type IGLApiResponse struct {
	Data []struct {
		Team struct {
			ID        string        `json:"_id"`
			Active    bool          `json:"active"`
			Dq        bool          `json:"dq"`
			Logo      string        `json:"logo"`
			Losses    int           `json:"losses"`
			PlayerIds []string      `json:"playerIds"`
			Schedule  []interface{} `json:"schedule"`
			Status    string        `json:"status"`
			TeamLimit int           `json:"teamLimit"`
			Wins      int           `json:"wins"`
			CaptainID string        `json:"captainId"`
			CircuitID string        `json:"circuitId"`
			Tier      float64       `json:tier,omitempty`
			Div       float64       `json:div,omitempty`
			CreatedAt time.Time     `json:"createdAt"`
			UpdatedAt time.Time     `json:"updatedAt"`
			V         int           `json:"__v"`
			Invite    struct {
				Token     string    `json:"token"`
				CreatedAt time.Time `json:"createdAt"`
				UpdatedAt time.Time `json:"updatedAt"`
			} `json:"invite"`
			FormattedName string `json:"formattedName"`
			Name          string `json:"name"`
		} `json:"team"`
		Player struct {
		} `json:"player"`
		Stats struct {
			MatchesWon  string `json:"Matches Won"`
			MatchesLost string `json:"Matches Lost"`
			GameWin     string `json:"Game Win %"`
			GamesWon    string `json:"Games Won"`
			GamesLost   string `json:"Games Lost"`
			MapWin      string `json:"Map Win %"`
			Kills       string `json:"Kills"`
			Deaths      string `json:"Deaths"`
			Berries     string `json:"Berries"`
			Snail       string `json:"Snail"`
		} `json:"stats"`
	} `json:"data"`
}

// OLD JSON parser
// results := result["data"].([]interface{})
// teams := []Team{}
// for _, v := range results {
// 	t := v.(map[string]interface{})["team"]
// 	s := v.(map[string]interface{})["stats"]
// 	team := Team{}
// 	team.Name = t.(map[string]interface{})["formattedName"].(string)
// 	team.Tier = int(t.(map[string]interface{})["tier"].(float64))
// 	team.Div = int(t.(map[string]interface{})["div"].(float64))
// 	if t.(map[string]interface{})["logo"] != nil {
// 		team.Img = t.(map[string]interface{})["logo"].(string)
// 	} else {
// 		team.Img = KQBAvatarImage
// 	}
// 	team.Stats.GamesWon, _ = strconv.Atoi(s.(map[string]interface{})["Games Won"].(string))
// 	team.Stats.GamesLost, _ = strconv.Atoi(s.(map[string]interface{})["Games Lost"].(string))
// 	team.Stats.MatchesWon, _ = strconv.Atoi(s.(map[string]interface{})["Matches Won"].(string))
// 	team.Stats.MatchesLost, _ = strconv.Atoi(s.(map[string]interface{})["Matches Lost"].(string))
// 	teams = append(teams, team)
// }
