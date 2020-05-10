package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"

	"github.com/manifoldco/promptui"
)

const (
	IglAPIURL = "https://indy-gaming-league-api.herokuapp.com/api/circuits/"
	IglWestID = "5e4c6b5178d46abdfeb49e71"
	IglEastID = "5e4b295560f132acbb31b8f5"
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
	var wg sync.WaitGroup
	wg.Add(1)
	StartHTTPServer()
	SetupCloseHandler()
	apiUrl, _ := PromptIglCircuit()
	teams := GetTeamInfo(apiUrl)
	home, err := PromptTeam("Home", teams)
	if err != nil {
		os.Exit(0)
	}
	away, err := PromptTeam("Away", teams)
	if err != nil {
		os.Exit(0)
	}
	s := Scoreboard{&home, &away, 0, 0, 0, 0}
	UpdateScoreBoard(&s)
	for s.HomeGames < 3 && s.AwayGames < 3 {
		err := RecordMapScore(&s)
		// Allow exit out if prompt errors out
		if err != nil {
			os.Exit(0)
		}
		UpdateScoreBoard(&s)
	}
	wg.Wait()
}

func GetTeamInfo(url string) []Team {
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
		}
		test.Stats.GamesWon, _ = strconv.Atoi(s.(map[string]interface{})["Games Won"].(string))
		test.Stats.GamesLost, _ = strconv.Atoi(s.(map[string]interface{})["Games Lost"].(string))
		test.Stats.MatchesWon, _ = strconv.Atoi(s.(map[string]interface{})["Matches Won"].(string))
		test.Stats.MatchesLost, _ = strconv.Atoi(s.(map[string]interface{})["Matches Lost"].(string))
		teams = append(teams, test)
	}
	// fmt.Println(teams)
	return teams
}

func PromptIglCircuit() (string, error) {
	prompt := promptui.Select{
		Label: fmt.Sprintf("Select IGL Circuit:"),
		Items: []string{"KQB East", "KQB West"},
	}

	i, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "", errors.New("Invalid Team Selection")
	}
	var url string

	if i == 0 {
		url = fmt.Sprintf("%s%s/results?bucket=igl-teamlogopics", IglAPIURL, IglEastID)
	} else {
		url = fmt.Sprintf("%s%s/results?bucket=igl-teamlogopics", IglAPIURL, IglWestID)
	}
	return url, nil
}

// PromptTeam
func PromptTeam(name string, teams []Team) (Team, error) {
	// Any type can be given to the select's item as long as the templates properly implement the dot notation
	// to display it.
	var color string
	if name == "Home" {
		color = "cyan"
	} else {
		color = "yellow"
	}
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   fmt.Sprintf("\U0001F41D {{ .Name | %s }}", color),
		Inactive: fmt.Sprintf("  {{ .Name | %s }}", color),
		Selected: fmt.Sprintf("\U0001F41D {{ .Name | %s }}", color),
		Details: `
	--------- Team ----------
	{{ "Name:" | faint }}	{{ .Name }}
	{{ "Tier:" | faint }}	{{ .Tier }}
	{{ "Division:" | faint }}	{{ .Div }}`,
	}

	prompt := promptui.Select{
		Label:     fmt.Sprintf("Select %s Team", name),
		Items:     teams,
		Templates: templates,
		Size:      6,
	}

	i, _, err := prompt.Run()

	if err != nil {
		log.Printf("Prompt failed %v\n", err)
		return Team{}, errors.New("Map Score Prompt Failed")
	}
	return teams[i], nil
}

// RecordMapScore
func RecordMapScore(s *Scoreboard) error {
	// Any type can be given to the select's item as long as the templates properly implement the dot notation
	// to display it.
	type MapWinPrompt struct {
		Option string
		Color  string
		ID     int
	}
	m := []MapWinPrompt{
		{fmt.Sprintf("%s (Home) Won Map", s.Home.Name), "cyan", 0},
		{fmt.Sprintf("%s (Away) Won Map", s.Away.Name), "yellow", 1},
	}
	templates := &promptui.SelectTemplates{
		Label:    "{{ . | red | bold }}?",
		Active:   "\U0001F525 {{ .Option | red | bold }}",
		Inactive: "  {{ .Option | faint }}",
		Selected: "\U0001F525 {{ .Option | red | bold }}",
	}

	prompt := promptui.Select{
		Label:     fmt.Sprintf("Select Who Won Map"),
		Items:     m,
		Templates: templates,
		Size:      2,
	}

	i, _, err := prompt.Run()

	if m[i].ID == 0 {
		s.IncrementHome()
		fmt.Printf("%s Score: %d Games %d Maps\n", s.Home.Name, s.HomeGames, s.HomeMaps)
	} else if m[i].ID == 1 {
		s.IncrementAway()
		fmt.Printf("%s: %d Games %d Maps\n", s.Away.Name, s.AwayGames, s.AwayMaps)
	}

	if err != nil {
		log.Printf("Map Score Prompt failed %v\n", err)
		return errors.New("Map Score Prompt Failed")
	}
	return nil
}

func SetupCloseHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		os.Exit(0)
	}()
}
