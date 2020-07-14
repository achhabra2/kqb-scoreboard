package main

import (
	"fmt"
	"log"
	"net/url"
	"strconv"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func GameType(a fyne.App) *fyne.Window {
	label := widget.NewLabel("Select Match Type")
	label.TextStyle.Bold = true
	label.Alignment = fyne.TextAlignCenter
	win := a.NewWindow("KQB Scoreboard App")
	IGLButton := widget.NewButton("IGL Match", func() {
		log.Println("Selected IGL KQB Scoreboard")
		content := IGLCircuitSelect(win)
		win.Resize(fyne.NewSize(400, 500))
		win.SetContent(content)
	})
	CustomButton := widget.NewButton("Custom Match", func() {
		log.Println("Selected Custom Match Type")
		content := CustomTeamSelection(win)
		win.Resize(fyne.NewSize(400, 500))
		win.SetContent(content)
	})
	container := fyne.NewContainerWithLayout(layout.NewVBoxLayout(),
		label, IGLButton, CustomButton)

	win.SetContent(container)
	win.Resize(fyne.NewSize(400, 500))
	win.Show()
	return &win
}

func CustomTeamSelection(w fyne.Window) *fyne.Container {
	label := widget.NewLabel("Enter Team Info")
	label.Alignment = fyne.TextAlignCenter
	label.TextStyle.Bold = true
	blueInput := widget.NewEntry()
	blueInput.SetPlaceHolder("Blue Team")
	goldInput := widget.NewEntry()
	goldInput.SetPlaceHolder("Gold Team")
	// win := a.NewWindow("Enter Team Info")

	options := make([]string, 10)
	for i := 0; i < 10; i++ {
		options[i] = strconv.Itoa(i)
	}
	var (
		blueWinsInt int
		goldWinsInt int
		blueLossInt int
		goldLossInt int
	)
	blueWins := widget.NewSelect(options, func(value string) {
		log.Println("Select set to", value)
		blueWinsInt, _ = strconv.Atoi(value)
	})
	blueLoss := widget.NewSelect(options, func(value string) {
		log.Println("Select set to", value)
		blueLossInt, _ = strconv.Atoi(value)
	})
	blueStandingsLabel := widget.NewLabel("Blue Wins-Losses")
	goldWins := widget.NewSelect(options, func(value string) {
		log.Println("Select set to", value)
		goldWinsInt, _ = strconv.Atoi(value)
	})
	goldLoss := widget.NewSelect(options, func(value string) {
		log.Println("Select set to", value)
		goldLossInt, _ = strconv.Atoi(value)
	})
	goldStandingsLabel := widget.NewLabel("Blue Wins-Losses")
	blueWLContainer := fyne.NewContainerWithLayout(layout.NewHBoxLayout(), blueStandingsLabel, blueWins, blueLoss)
	goldWLContainer := fyne.NewContainerWithLayout(layout.NewHBoxLayout(), goldStandingsLabel, goldWins, goldLoss)

	saveButton := widget.NewButton("Start Scoreboard", func() {
		log.Println("Scoreboard Starting")
		blueTeam := Team{blueInput.Text, "avatar.png", 1, 1, Stats{0, 0, blueWinsInt, blueLossInt}}
		goldTeam := Team{goldInput.Text, "avatar.png", 1, 1, Stats{0, 0, goldWinsInt, goldLossInt}}
		s = Scoreboard{&blueTeam, &goldTeam, 0, 0, 0, 0}
		StartScoreboard(w)
	})
	saveButton.Style = widget.PrimaryButton

	container := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), label, blueInput, blueWLContainer, goldInput, goldWLContainer, saveButton)
	return container
}

func IGLTeamSelection(w fyne.Window, apiUrl string) *fyne.Container {
	label := widget.NewLabel("Select Teams")
	label.Alignment = fyne.TextAlignCenter
	label.TextStyle.Bold = true
	blueLabel := widget.NewLabel("Blue Team")
	goldLabel := widget.NewLabel("Gold Team")

	ch := make(chan []Team)

	w.SetContent(ProgressIndicator())
	go GetTeamInfo(apiUrl, ch)
	teams := <-ch

	options := make([]string, len(teams))
	var blueTeam Team
	var goldTeam Team

	saveButton := widget.NewButton("Start Scoreboard", func() {
		log.Println("Saved.")
		s = Scoreboard{&blueTeam, &goldTeam, 0, 0, 0, 0}
		UpdateTeamLogo(&blueTeam)
		UpdateTeamLogo(&goldTeam)
		StartScoreboard(w)
	})

	saveButton.Style = widget.PrimaryButton
	saveButton.Disable()

	for i, team := range teams {
		options[i] = team.Name
	}
	blueTeamSelect := widget.NewSelect(options, func(value string) {
		for i, v := range options {
			if v == value {
				blueTeam = teams[i]
			}
		}
		log.Println("Select set to", value)
		if blueTeam.Name != "" && goldTeam.Name != "" {
			saveButton.Enable()
		}
	})
	goldTeamSelect := widget.NewSelect(options, func(value string) {
		for i, v := range options {
			if v == value {
				goldTeam = teams[i]
			}
		}
		log.Println("Select set to", value)
		if blueTeam.Name != "" && goldTeam.Name != "" {
			saveButton.Enable()
		}
	})

	blueTeamContainer := fyne.NewContainerWithLayout(layout.NewHBoxLayout(), layout.NewSpacer(), blueLabel, blueTeamSelect, layout.NewSpacer())
	goldTeamContainer := fyne.NewContainerWithLayout(layout.NewHBoxLayout(), layout.NewSpacer(), goldLabel, goldTeamSelect, layout.NewSpacer())

	container := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), label, blueTeamContainer, goldTeamContainer, saveButton)

	return container
}

func IGLCircuitSelect(w fyne.Window) *fyne.Container {
	label := widget.NewLabel("Select Circuit")
	label.Alignment = fyne.TextAlignCenter
	label.TextStyle.Bold = true

	ch := make(chan []IGLCircuit)

	w.SetContent(ProgressIndicator())
	go GetIGLCircuits(ch)
	circuits := <-ch

	kqbCircuits := []IGLCircuit{}

	for _, circuit := range circuits {
		if circuit.Game == "KILLER QUEEN BLACK" {
			kqbCircuits = append(kqbCircuits, circuit)
		}
	}

	options := make([]string, len(kqbCircuits))

	for i, circuit := range kqbCircuits {
		options[i] = circuit.Region + " " + circuit.Game
	}

	var url string

	nextButton := widget.NewButton("Next", func() {
		log.Println("Saved.")
		w.SetContent(IGLTeamSelection(w, url))
	})

	nextButton.Style = widget.PrimaryButton
	nextButton.Disable()

	circuitSelect := widget.NewSelect(options, func(value string) {
		for i, option := range options {
			if option == value {
				url = fmt.Sprintf("%s%s/results?bucket=igl-teamlogopics", IglAPIURL, kqbCircuits[i].ID)
			}
		}

		nextButton.Enable()
	})

	container := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), label, circuitSelect, nextButton)

	return container
}

func ScoreboardContent(w fyne.Window) *fyne.Container {
	scoreboardLabel := widget.NewLabel("Scoreboard Controller")
	scoreboardLabel.Alignment = fyne.TextAlignCenter
	scoreboardLabel.TextStyle.Bold = true
	blueLabel := widget.NewLabel(fmt.Sprintf("%s (Blue)", s.Home.Name))
	blueLabel.Alignment = fyne.TextAlignCenter
	goldLabel := widget.NewLabel(fmt.Sprintf("%s (Gold)", s.Away.Name))
	goldLabel.Alignment = fyne.TextAlignCenter
	mapsLabel := widget.NewLabel("Maps: ")
	setsLabel := widget.NewLabel("Sets: ")
	blueMaps := widget.NewLabel("0")
	goldMaps := widget.NewLabel("0")
	blueSets := widget.NewLabel("0")
	goldSets := widget.NewLabel("0")
	blueScoreboard := fyne.NewContainerWithLayout(layout.NewFormLayout(), mapsLabel, blueMaps, setsLabel, blueSets)
	goldScoreboard := fyne.NewContainerWithLayout(layout.NewFormLayout(), mapsLabel, goldMaps, setsLabel, goldSets)
	scoreboardContainer := fyne.NewContainerWithLayout(layout.NewGridLayout(2), blueLabel, goldLabel, blueScoreboard, goldScoreboard)
	incrementBlue := widget.NewButtonWithIcon("Increment Blue", theme.ContentAddIcon(), func() {
		s.IncrementHome()
		UpdateScoreBoard(&s)
		blueMaps.Text = strconv.Itoa(s.HomeMaps)
		blueSets.Text = strconv.Itoa(s.HomeGames)
		goldMaps.Text = strconv.Itoa(s.AwayMaps)
		goldSets.Text = strconv.Itoa(s.AwayGames)
		scoreboardContainer.Refresh()
	})
	incrementGold := widget.NewButtonWithIcon("Increment Gold", theme.ContentAddIcon(), func() {
		s.IncrementAway()
		UpdateScoreBoard(&s)
		blueMaps.Text = strconv.Itoa(s.HomeMaps)
		blueSets.Text = strconv.Itoa(s.HomeGames)
		goldMaps.Text = strconv.Itoa(s.AwayMaps)
		goldSets.Text = strconv.Itoa(s.AwayGames)
		scoreboardContainer.Refresh()
	})
	decrementBlue := widget.NewButtonWithIcon("Decrement Blue", theme.ContentClearIcon(), func() {
		s.DecrementHome()
		UpdateScoreBoard(&s)
		blueMaps.Text = strconv.Itoa(s.HomeMaps)
		blueSets.Text = strconv.Itoa(s.HomeGames)
		goldMaps.Text = strconv.Itoa(s.AwayMaps)
		goldSets.Text = strconv.Itoa(s.AwayGames)
		scoreboardContainer.Refresh()
	})
	decrementGold := widget.NewButtonWithIcon("Decrement Gold", theme.ContentClearIcon(), func() {
		s.DecrementAway()
		UpdateScoreBoard(&s)
		blueMaps.Text = strconv.Itoa(s.HomeMaps)
		blueSets.Text = strconv.Itoa(s.HomeGames)
		goldMaps.Text = strconv.Itoa(s.AwayMaps)
		goldSets.Text = strconv.Itoa(s.AwayGames)
		scoreboardContainer.Refresh()
	})
	blueContainer := fyne.NewContainerWithLayout(layout.NewHBoxLayout(), layout.NewSpacer(), incrementBlue, incrementGold, layout.NewSpacer())
	goldContainer := fyne.NewContainerWithLayout(layout.NewHBoxLayout(), layout.NewSpacer(), decrementBlue, decrementGold, layout.NewSpacer())
	scoreboardURL, _ := url.Parse("http://localhost:8080")
	link := widget.NewHyperlink("Scoreboard", scoreboardURL)
	link.Alignment = fyne.TextAlignCenter

	resetButton := widget.NewButtonWithIcon("Reset", theme.DeleteIcon(), func() {
		s.HomeMaps = 0
		s.HomeGames = 0
		s.AwayGames = 0
		s.HomeGames = 0
		blueMaps.Text = strconv.Itoa(s.HomeMaps)
		blueSets.Text = strconv.Itoa(s.HomeGames)
		goldMaps.Text = strconv.Itoa(s.AwayMaps)
		goldSets.Text = strconv.Itoa(s.AwayGames)
		scoreboardContainer.Refresh()
		UpdateScoreBoard(&s)
	})
	resetButtonContainer := fyne.NewContainerWithLayout(layout.NewHBoxLayout(), layout.NewSpacer(), resetButton, layout.NewSpacer())
	starTimerButton := widget.NewButton("Start Timer", func() {
		UpdateTimer("StartTimer")
	})
	stopTimerButton := widget.NewButton("Stop Timer", func() {
		UpdateTimer("StopTimer")
	})
	resetTimerButton := widget.NewButton("Reset Timer", func() {
		UpdateTimer("ResetTimer")
	})
	hideTimerButton := widget.NewButton("Show/Hide Timer", func() {
		UpdateTimer("ToggleTimer")
	})
	timerContainer := fyne.NewContainerWithLayout(layout.NewHBoxLayout(), layout.NewSpacer(), starTimerButton, stopTimerButton, resetTimerButton, hideTimerButton, layout.NewSpacer())
	container := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), scoreboardLabel, blueContainer, goldContainer, scoreboardContainer, resetButtonContainer, link, timerContainer)
	return container
}

func StartScoreboard(w fyne.Window) {
	w.Resize(fyne.NewSize(400, 500))
	w.SetContent(ScoreboardContent(w))
}

func ProgressIndicator() *fyne.Container {
	label := widget.NewLabel("Fetching data...")
	label.Alignment = fyne.TextAlignCenter
	label.TextStyle.Bold = true
	infinite := widget.NewProgressBarInfinite()
	container := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), label, infinite)
	return container
}
