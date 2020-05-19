package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

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
	setupLogs()
	// var wg sync.WaitGroup
	// wg.Add(1)
	StartHTTPServer()
	// SetupCloseHandler()
	myApp := app.New()
	_ = GameType(myApp)
	myApp.Run()
	// RunMatch()
	// wg.Wait()
	tidyUp()
}

func GetTeamInfo(url string, c chan []Team) {
	fmt.Println("Fetching Team Information from IGL...")
	resp, err := http.Get(url)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	// jsonFile, err := os.Open("response.json")
	// // if we os.Open returns an error then handle it
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("Successfully Opened users.json")
	// // defer the closing of our jsonFile so that we can parse it later on
	// defer jsonFile.Close()

	// byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(body), &result)
	results := result["data"].([]interface{})
	teams := []Team{}
	for _, v := range results {
		t := v.(map[string]interface{})["team"]
		s := v.(map[string]interface{})["stats"]
		test := Team{}
		test.Name = t.(map[string]interface{})["formattedName"].(string)
		test.Tier = int(t.(map[string]interface{})["tier"].(float64))
		test.Div = int(t.(map[string]interface{})["div"].(float64))
		if t.(map[string]interface{})["logo"] != nil {
			test.Img = t.(map[string]interface{})["logo"].(string)
		} else {
			test.Img = KQBAvatarImage
		}
		test.Stats.GamesWon, _ = strconv.Atoi(s.(map[string]interface{})["Games Won"].(string))
		test.Stats.GamesLost, _ = strconv.Atoi(s.(map[string]interface{})["Games Lost"].(string))
		test.Stats.MatchesWon, _ = strconv.Atoi(s.(map[string]interface{})["Matches Won"].(string))
		test.Stats.MatchesLost, _ = strconv.Atoi(s.(map[string]interface{})["Matches Lost"].(string))
		teams = append(teams, test)
	}
	c <- teams
}
