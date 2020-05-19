package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const CircuitAPIURL = "https://indy-gaming-league-api.herokuapp.com/api/circuits?data={%22motherCircuitId%22:%225e4b290a420594ace7e97726%22,%22live%22:false,%22active%22:true,%22hidden%22:false,%20%22game%22:%22KILLER%20QUEEN%20BLACK%22}"

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