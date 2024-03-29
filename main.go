package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

const (
	PlaceholderImage = "https://bulma.io/images/placeholders/128x128.png"
	KQBAvatarImage   = "avatar.png"
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
var s Scoreboard = Scoreboard{&h, &a, 0, 0, 0, 0, []ScoreboardSet{}}
var logoPath string
var themePath string
var themes []string
var selectedTheme string
var FyneApp fyne.App

func setupLogs() {
	f, err := os.OpenFile("./scoreboard-output.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
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

	// Create Directory For Logos
	SetupLogoDirectory()

	// Setup HTTP handler and websocket handler
	StartHTTPServer()

	// Create Fyne App
	FyneApp = app.New()
	FyneApp.SetIcon(resourceIconPng)
	_ = GameType(FyneApp)

	// Add Event Hotkey Bindings
	go AddEventHotkeys()
	FyneApp.Run()

	// Called before application exist
	tidyUp()
}

// SetupLogoDirectory creates a directory to cache the local logo files
func SetupLogoDirectory() {

	// Get Current Working Directory
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)

	// Set Logo Path
	logoPath = filepath.Join(path, "logo")
	fmt.Println(logoPath)

	// Set Theme Path
	themePath = filepath.Join(path, "themes")

	// Create directory if it doesn't exist
	if _, err := os.Stat(logoPath); os.IsNotExist(err) {
		os.Mkdir(logoPath, 0755)
	}

	// Create directory if it doesn't exist
	if _, err := os.Stat(themePath); os.IsNotExist(err) {
		themePath = ""
	}

	files, err := ioutil.ReadDir(themePath)
	if err != nil {
		log.Println(err)
	}

	themes = append(themes, "default", "bgl")

	for _, file := range files {
		name := file.Name()
		themes = append(themes, name)
	}

}

func tidyUp() {

	// Remove logopath directory and files
	os.RemoveAll(logoPath)
	os.Remove("./output.log")
	log.Println("Exiting...")
}
