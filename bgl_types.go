package main

import "time"

type MatchResponse struct {
	Count    int         `json:"count"`
	Next     interface{} `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		ID   int `json:"id"`
		Home struct {
			ID      int    `json:"id"`
			Name    string `json:"name"`
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
			Group  interface{} `json:"group"`
			Wins   int         `json:"wins"`
			Losses int         `json:"losses"`
		} `json:"home"`
		Away struct {
			ID      int    `json:"id"`
			Name    string `json:"name"`
			Members []struct {
				ID              int         `json:"id"`
				Name            string      `json:"name"`
				NamePhonetic    string      `json:"name_phonetic"`
				Pronouns        string      `json:"pronouns"`
				DiscordUsername string      `json:"discord_username"`
				TwitchUsername  string      `json:"twitch_username"`
				Bio             interface{} `json:"bio"`
				Emoji           interface{} `json:"emoji"`
				AvatarURL       string      `json:"avatar_url"`
				Modified        time.Time   `json:"modified"`
				Created         time.Time   `json:"created"`
			} `json:"members"`
			Group  interface{} `json:"group"`
			Wins   int         `json:"wins"`
			Losses int         `json:"losses"`
		} `json:"away"`
		Circuit struct {
			ID     int `json:"id"`
			Season struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"season"`
			Region      string `json:"region"`
			Tier        string `json:"tier"`
			Name        string `json:"name"`
			VerboseName string `json:"verbose_name"`
		} `json:"circuit"`
		CircuitDisplay string `json:"circuit_display"`
		Round          struct {
			Number string `json:"number"`
			Name   string `json:"name"`
		} `json:"round"`
		StartTime     time.Time `json:"start_time"`
		TimeUntil     string    `json:"time_until"`
		Scheduled     bool      `json:"scheduled"`
		PrimaryCaster struct {
			ID         int         `json:"id"`
			Name       string      `json:"name"`
			BioLink    interface{} `json:"bio_link"`
			StreamLink string      `json:"stream_link"`
		} `json:"primary_caster"`
		SecondaryCasters []interface{} `json:"secondary_casters"`
		Result           interface{}   `json:"result"`
		VodLink          interface{}   `json:"vod_link"`
	} `json:"results"`
}

type TeamsResponse struct {
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
