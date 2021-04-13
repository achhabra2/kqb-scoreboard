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
	"time"
)

type TeamsResult struct {
	Count    int         `json:"count"`
	Next     interface{} `json:"next"`
	Previous interface{} `json:"previous"`
	Teams    []struct {
		ID            int         `json:"id"`
		Name          string      `json:"name"`
		Circuit       int         `json:"circuit"`
		Group         interface{} `json:"group"`
		IsActive      bool        `json:"is_active"`
		CanAddMembers bool        `json:"can_add_members"`
		Dynasty       interface{} `json:"dynasty"`
		Captain       struct {
			ID              int         `json:"id"`
			Name            string      `json:"name"`
			NamePhonetic    interface{} `json:"name_phonetic"`
			Pronouns        string      `json:"pronouns"`
			DiscordUsername interface{} `json:"discord_username"`
			TwitchUsername  string      `json:"twitch_username"`
			Bio             string      `json:"bio"`
			Emoji           interface{} `json:"emoji"`
			AvatarURL       string      `json:"avatar_url"`
			Modified        time.Time   `json:"modified"`
			Created         time.Time   `json:"created"`
		} `json:"captain"`
		Members []struct {
			ID              int         `json:"id"`
			Name            string      `json:"name"`
			NamePhonetic    string      `json:"name_phonetic"`
			Pronouns        string      `json:"pronouns"`
			DiscordUsername string      `json:"discord_username"`
			TwitchUsername  string      `json:"twitch_username"`
			Bio             string      `json:"bio"`
			Emoji           interface{} `json:"emoji"`
			AvatarURL       string      `json:"avatar_url"`
			Modified        time.Time   `json:"modified"`
			Created         time.Time   `json:"created"`
		} `json:"members"`
		Modified time.Time `json:"modified"`
		Created  time.Time `json:"created"`
		Wins     int       `json:"wins"`
		Losses   int       `json:"losses"`
	} `json:"results"`
}

func GetBGLTeams() TeamsResult {

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

	var teamsResult TeamsResult
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
