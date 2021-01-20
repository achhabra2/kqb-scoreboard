package main

import "log"

// Scoreboard defines the data structure for the home team and away team
type Scoreboard struct {
	Home      *Team
	Away      *Team
	HomeMaps  int
	AwayMaps  int
	HomeGames int
	AwayGames int
	//	BlueBerries int
	//	GoldBerries int
	//	BlueEggs    int
	//	GoldEggs    int
	Sets []ScoreboardSet
}

type ScoreboardSet struct {
	Away int
	Home int
}

// IncrementHome accounts for when the home team wins a map
func (s *Scoreboard) IncrementHome() {
	s.HomeMaps++
	if s.HomeMaps == 3 {
		s.Sets = append(s.Sets, ScoreboardSet{Away: s.AwayMaps, Home: s.HomeMaps})
		s.HomeMaps = 0
		s.AwayMaps = 0
		s.HomeGames++
	}
	log.Printf("Scoreboard Update: %s Won Map\n", s.Home.Name)
	log.Printf("%s: %d Games %d Maps\n", s.Home.Name, s.HomeGames, s.HomeMaps)
}

// DecrementHome accounts for an error in scoring
func (s *Scoreboard) DecrementHome() {
	if s.HomeMaps > 0 {
		s.HomeMaps--
	}
}

// IncrementAway accounts for when the away team wins a map
func (s *Scoreboard) IncrementAway() {
	s.AwayMaps++
	if s.AwayMaps == 3 {
		s.Sets = append(s.Sets, ScoreboardSet{Away: s.AwayMaps, Home: s.HomeMaps})
		s.HomeMaps = 0
		s.AwayMaps = 0
		s.AwayGames++
	}
	log.Printf("Scoreboard Update: %s Won Map\n", s.Away.Name)
	log.Printf("%s: %d Games %d Maps\n", s.Away.Name, s.AwayGames, s.AwayMaps)
}

// DecrementAway accounts for an error in scoring
func (s *Scoreboard) DecrementAway() {
	if s.AwayMaps > 0 {
		s.AwayMaps--
	}
}
