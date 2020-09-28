package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const CircuitAPIURL = "https://indy-gaming-league-api.herokuapp.com/api/circuits?data=%7B%22motherCircuitId%22:%225f380956d98a1a47bb52b692%22,%22live%22:false,%22active%22:true,%22hidden%22:false,%22game%22:%22KILLER%20QUEEN%20BLACK%22%7D"

type IGLCircuit struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Region string `json:"region"`
	Game   string `json:"game"`
}

func (c IGLCircuit) String() string {
	return fmt.Sprintf("Circuit ID: %s, Name: %s, Region: %s\n", c.ID, c.Name, c.Region)
}

type IGLCircuitData struct {
	Circuits []IGLCircuit `json:"circuits"`
}

func GetIGLCircuits(c chan []IGLCircuit) {
	fmt.Println("Fetching Circuit Information from IGL...")
	resp, err := http.Get(CircuitAPIURL)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var result IGLCircuitData
	json.Unmarshal([]byte(body), &result)

	// fmt.Println(result.Circuits)

	c <- result.Circuits
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

// UpdateTeamLogo will Take Team as an input and download the logo locally
// as a png file
func UpdateTeamLogo(t *Team) {

	if t.Img == "/avatar.png" {
		return
	}

	log.Printf("Downloading Logo %s\n", t.Img)
	logoFile := fmt.Sprintf("%s.png", strings.Replace(t.Name, " ", "", -1))
	err := DownloadFile(filepath.Join(logoPath, logoFile), t.Img)
	if err != nil {
		log.Println("Could not save logo image")
		log.Println(err)
		return
	}
	t.Img = "http://localhost:8080/logo/" + logoFile
}
