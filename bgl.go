// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    matchResult, err := UnmarshalMatchResult(bytes)
//    bytes, err = matchResult.Marshal()

package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func GetBGLTeams() TeamsResponse {

	url := "https://api.beegame.gg/teams/?is_active=true"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		log.Println(err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}

	var teamsResult TeamsResponse
	err = json.Unmarshal(body, &teamsResult)
	if err != nil {
		log.Println(err)
	}

	return teamsResult

}

func GetTeamInfo(c chan []Team) {
	log.Println("Fetching Team Information from BGL...")
	bglTeams := GetBGLTeams()
	teams := []Team{}
	for _, v := range bglTeams.Teams {
		team := Team{}

		team.Name = v.Name
		team.Tier = v.Circuit
		team.Img = KQBAvatarImage

		team.Stats.MatchesWon = v.Wins
		team.Stats.MatchesLost = v.Losses
		teams = append(teams, team)
	}

	c <- teams
}

func GetBGLMatches() MatchResponse {
	url := "https://api.beegame.gg/matches/?days=7&scheduled=true&limit=100"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		log.Println(err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}

	var matchResponse MatchResponse
	err = json.Unmarshal(body, &matchResponse)
	if err != nil {
		log.Println(err)
	}

	return matchResponse
}

func MatchesToTeamMap(matches MatchResponse) MatchMap {
	matchMap := make(MatchMap)
	
	for _, match := range matches.Results {
		title := match.Away.Name + " vs " + match.Home.Name
		home := Team{
			Name: match.Home.Name,
			Img: KQBAvatarImage,
			Stats: Stats{
				MatchesWon:match.Home.Wins,
				MatchesLost: match.Home.Losses,
			},
		}

		away := Team{
			Name: match.Away.Name,
			Img: KQBAvatarImage,
			Stats: Stats{
				MatchesWon:match.Away.Wins,
				MatchesLost: match.Away.Losses,
			},
		}

		matchMap[title] = []Team{away, home}
	}

	return matchMap
}

func GetMatchInfo(c chan MatchMap) {
	log.Println("Fetching Match Information from BGL...")
	matches := GetBGLMatches()
	matchMap := MatchesToTeamMap(matches)

	c <- matchMap
}

type MatchMap map[string][]Team