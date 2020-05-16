package main

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/manifoldco/promptui"
)

func PromptIGLCustom() (result string) {

	prompt := promptui.Select{
		Label: "Is this an IGL Match or Custom Match?",
		Items: []string{"IGL Match", "Custom Match"},
	}

	idx, result, err := prompt.Run()

	if err != nil {
		log.Printf("Prompt failed %v\n", err)
		return ""
	}

	if idx == 0 {
		result = "IGL"
	} else {
		result = "Custom"
	}

	return
}

func CustomFlow() Scoreboard {
	blue := PromptCustomName("What is the Blue Team Name? ")
	gold := PromptCustomName("What is the Gold Team Name? ")

	standing, err := regexp.Compile(`^(\d+)\-(\d+)$`)
	if err != nil {
		log.Println("Could not compile regex", err)
	}

	blueStanding := PromptCustomStanding("Blue")
	blueWon := standing.FindAllSubmatch([]byte(blueStanding), -1)
	goldStanding := PromptCustomStanding("Gold")
	goldWon := standing.FindAllSubmatch([]byte(goldStanding), -1)

	log.Println("Regex matching: ")
	log.Printf("%q\n", blueWon)
	log.Printf("%q\n", goldWon)

	blueMatchesWon, _ := strconv.Atoi(string(blueWon[0][1]))
	blueMatchesLost, _ := strconv.Atoi(string(blueWon[0][2]))

	goldMatchesWon, _ := strconv.Atoi(string(goldWon[0][1]))
	goldMatchesLost, _ := strconv.Atoi(string(goldWon[0][2]))

	blueTeam := Team{blue, "avatar.png", 1, 1, Stats{0, 0, blueMatchesWon, blueMatchesLost}}
	goldTeam := Team{gold, "avatar.png", 1, 1, Stats{0, 0, goldMatchesWon, goldMatchesLost}}

	return Scoreboard{&blueTeam, &goldTeam, 0, 0, 0, 0}

}

func PromptCustomName(name string) (result string) {

	prompt := promptui.Prompt{
		Label:     name,
		AllowEdit: true,
	}

	result, err := prompt.Run()

	if err != nil {
		log.Printf("Prompt failed %v\n", err)
		return
	}

	log.Printf("You choose %q\n", result)
	return
}

func PromptCustomStanding(team string) (result string) {
	validate := func(input string) error {
		match, _ := regexp.Match(`^\d+\-\d+$`, []byte(input))
		if !match {
			return errors.New("invalid standing win/loss entry")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:     fmt.Sprintf("What is the %s Team Matches Won-Lost?", team),
		Default:   "0-0",
		AllowEdit: true,
		Validate:  validate,
	}

	result, err := prompt.Run()

	if err != nil {
		log.Printf("Prompt failed %v\n", err)
		return
	}

	log.Printf("You choose %q\n", result)
	return
}
