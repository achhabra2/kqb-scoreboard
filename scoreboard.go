package main

// Scoreboard defines the data structure for the home team and away team
type Scoreboard struct {
	Home      *Team
	Away      *Team
	HomeMaps  int
	AwayMaps  int
	HomeGames int
	AwayGames int
}

// IncrementHome accounts for when the home team wins a map
func (s *Scoreboard) IncrementHome() {
	s.HomeMaps++
	if s.HomeMaps == 3 {
		s.HomeMaps = 0
		s.AwayMaps = 0
		s.HomeGames++
	}
	// fmt.Printf("Scoreboard Update: %s Won Map\n", s.home.Name)
	// fmt.Printf("%s: %d Games %d Maps\n", s.home.Name, s.homeGames, s.homeMaps)
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
		s.HomeMaps = 0
		s.AwayMaps = 0
		s.AwayGames++
	}
	// fmt.Printf("Scoreboard Update: %s Won Map\n", s.away.Name)
	// fmt.Printf("%s: %d Games %d Maps\n", s.away.Name, s.awayGames, s.awayMaps)
}

// DecrementAway accounts for an error in scoring
func (s *Scoreboard) DecrementAway() {
	if s.AwayMaps > 0 {
		s.AwayMaps--
	}
}
