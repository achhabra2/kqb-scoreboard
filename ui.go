package main

import (
	"fmt"
	"log"
	"net/url"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/aquilax/truncate"
)

func GameType(a fyne.App) *fyne.Window {
	label := widget.NewLabel("Select Match Type")
	label.TextStyle.Bold = true
	label.Alignment = fyne.TextAlignCenter
	win := a.NewWindow("KQB Scoreboard App")
	IGLButton := widget.NewButton("BGL Match", func() {
		log.Println("Selected BGL KQB Scoreboard")
		content := BGLMatchSelection(win)
		win.Resize(fyne.NewSize(660, 500))
		win.SetContent(content)
	})
	CustomButton := widget.NewButton("Custom Match", func() {
		log.Println("Selected Custom Match Type")
		content := CustomTeamSelection(win)
		win.Resize(fyne.NewSize(660, 500))
		win.SetContent(content)
	})
	cont := container.NewVBox(label, IGLButton, CustomButton)

	win.SetContent(cont)
	win.Resize(fyne.NewSize(660, 500))
	win.CenterOnScreen()
	win.Show()
	go CheckForUpdates(win)
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
	blueWLContainer := container.NewHBox(blueStandingsLabel, blueWins, blueLoss)
	goldWLContainer := container.NewHBox(goldStandingsLabel, goldWins, goldLoss)

	saveButton := widget.NewButton("Start Scoreboard", func() {
		log.Println("Scoreboard Starting")
		blueTeam := Team{blueInput.Text, "avatar.png", 1, 1, Stats{0, 0, blueWinsInt, blueLossInt}}
		goldTeam := Team{goldInput.Text, "avatar.png", 1, 1, Stats{0, 0, goldWinsInt, goldLossInt}}
		s = Scoreboard{&blueTeam, &goldTeam, 0, 0, 0, 0, []ScoreboardSet{}}
		StartScoreboard(w)
	})
	saveButton.Importance = widget.HighImportance

	cont := container.NewVBox(label, blueInput, blueWLContainer, goldInput, goldWLContainer, saveButton)
	return cont
}

func BGLMatchSelection(w fyne.Window) *fyne.Container {
	label := widget.NewLabel("Select Match Info")
	label.Alignment = fyne.TextAlignCenter
	label.TextStyle.Bold = true
	blueLabel := widget.NewLabel("Blue Team")
	goldLabel := widget.NewLabel("Gold Team")

	blueTeamSelect := widget.NewSelect([]string{"Blue"}, func(val string) {})
	goldTeamSelect := widget.NewSelect([]string{"Gold"}, func(val string) {})
	blueTeamContainer := container.NewHBox(layout.NewSpacer(), blueLabel, blueTeamSelect, layout.NewSpacer())
	goldTeamContainer := container.NewHBox(layout.NewSpacer(), goldLabel, goldTeamSelect, layout.NewSpacer())

	themeSelectLabel := widget.NewLabel("Select Theme")
	themeSelect := widget.NewSelect(themes, func(theme string) {
		selectedTheme = theme
	})
	themeSelectContainer := container.NewHBox(layout.NewSpacer(), themeSelectLabel, themeSelect, layout.NewSpacer())
	// blueTeamSelect.Resize(fyne.NewSize(500, 100))
	blueTeamContainer.Hide()
	goldTeamContainer.Hide()
	// ch := make(chan []Team)
	ch := make(chan MatchMap)
	w.SetContent(ProgressIndicator())
	go GetMatchInfo(ch)
	matchMap := <-ch

	var blueTeam Team
	var goldTeam Team

	saveButton := widget.NewButton("Start Scoreboard", func() {
		log.Println("Saved.")
		s = Scoreboard{&blueTeam, &goldTeam, 0, 0, 0, 0, []ScoreboardSet{}}
		UpdateTeamLogo(&blueTeam)
		UpdateTeamLogo(&goldTeam)
		StartScoreboard(w)
		UpdateStaticRoute()
	})

	saveButton.Importance = widget.HighImportance
	saveButton.Disable()

	matchOptions := make([]string, 0)
	for k, _ := range matchMap {
		matchOptions = append(matchOptions, k)
	}
	matchLabel := widget.NewLabel("Select Match")
	matchLabel.Alignment = fyne.TextAlignCenter

	matchSelect := widget.NewSelect(matchOptions, func(value string) {
		selectedMatchTeams := matchMap[value]
		teamOptions := []string{selectedMatchTeams[0].Name, selectedMatchTeams[1].Name}
		teamOptionsFormatted := make([]string, len(teamOptions))
		for idx, opt := range teamOptions {
			teamOptionsFormatted[idx] = formatTeamName(opt)
		}
		blueTeamSelect.Options = teamOptionsFormatted
		blueTeamSelect.OnChanged = func(val string) {
			idx := blueTeamSelect.SelectedIndex()
			blueTeam = selectedMatchTeams[idx]
			log.Println("Select blue team to", value)
			if blueTeam.Name != "" && goldTeam.Name != "" {
				saveButton.Enable()
			}
		}
		goldTeamSelect.Options = teamOptionsFormatted
		goldTeamSelect.OnChanged = func(val string) {
			idx := goldTeamSelect.SelectedIndex()
			goldTeam = selectedMatchTeams[idx]
			log.Println("Select gold team to", value)
			if blueTeam.Name != "" && goldTeam.Name != "" {
				saveButton.Enable()
			}
		}

		blueTeamContainer.Refresh()
		goldTeamContainer.Refresh()
		blueTeamContainer.Show()
		goldTeamContainer.Show()
	})

	matchSelectContainer := container.NewHBox(layout.NewSpacer(), matchLabel, matchSelect, layout.NewSpacer())

	container := container.NewVBox(label, matchSelectContainer, blueTeamContainer, goldTeamContainer, themeSelectContainer, saveButton)

	return container
}

func ScoreboardContent(w fyne.Window, SetupEventHooks func(func())) *fyne.Container {
	scoreboardLabel := widget.NewLabel("Scoreboard Controller")
	scoreboardLabel.Alignment = fyne.TextAlignCenter
	scoreboardLabel.TextStyle.Bold = true
	blueLabel := widget.NewLabel(fmt.Sprintf("%s (Blue)", formatTeamName(s.Home.Name)))
	blueLabel.Alignment = fyne.TextAlignCenter
	goldLabel := widget.NewLabel(fmt.Sprintf("%s (Gold)", formatTeamName(s.Away.Name)))
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
	RefreshScoreboardUI := func() {
		blueMaps.Text = strconv.Itoa(s.HomeMaps)
		blueSets.Text = strconv.Itoa(s.HomeGames)
		goldMaps.Text = strconv.Itoa(s.AwayMaps)
		goldSets.Text = strconv.Itoa(s.AwayGames)
		scoreboardContainer.Refresh()
	}
	SetupEventHooks(RefreshScoreboardUI)
	incrementBlue := widget.NewButtonWithIcon("Increment Blue", theme.ContentAddIcon(), func() {
		s.IncrementHome()
		UpdateScoreBoard(&s)
		RefreshScoreboardUI()
	})
	incrementGold := widget.NewButtonWithIcon("Increment Gold", theme.ContentAddIcon(), func() {
		s.IncrementAway()
		UpdateScoreBoard(&s)
		RefreshScoreboardUI()
	})
	decrementBlue := widget.NewButtonWithIcon("Decrement Blue", theme.ContentClearIcon(), func() {
		s.DecrementHome()
		UpdateScoreBoard(&s)
		RefreshScoreboardUI()
	})
	decrementGold := widget.NewButtonWithIcon("Decrement Gold", theme.ContentClearIcon(), func() {
		s.DecrementAway()
		UpdateScoreBoard(&s)
		RefreshScoreboardUI()
	})

	blueContainer := container.NewHBox(layout.NewSpacer(), incrementBlue, incrementGold, layout.NewSpacer())
	goldContainer := container.NewHBox(layout.NewSpacer(), decrementBlue, decrementGold, layout.NewSpacer())
	scoreboardURL, _ := url.Parse("http://localhost:8080/static/")
	link := widget.NewHyperlink("Scoreboard", scoreboardURL)
	link.Alignment = fyne.TextAlignCenter

	resetButton := widget.NewButtonWithIcon("Reset", theme.DeleteIcon(), func() {
		s.HomeMaps = 0
		s.HomeGames = 0
		s.AwayGames = 0
		s.AwayMaps = 0
		s.Sets = []ScoreboardSet{}
		RefreshScoreboardUI()
		UpdateScoreBoard(&s)
	})
	resetButtonContainer := container.NewHBox(layout.NewSpacer(), resetButton, layout.NewSpacer())
	// starTimerButton := widget.NewButton("Start Timer", func() {
	// 	UpdateTimer("StartTimer")
	// })
	// stopTimerButton := widget.NewButton("Stop Timer", func() {
	// 	UpdateTimer("StopTimer")
	// })
	// resetTimerButton := widget.NewButton("Reset Timer", func() {
	// 	UpdateTimer("ResetTimer")
	// })
	// hideTimerButton := widget.NewButton("Show/Hide Timer", func() {
	// 	UpdateTimer("ToggleTimer")
	// })
	aboutButton := widget.NewButton("About", func() {
		ShowAboutWindow()
	})
	aboutContainer := container.NewHBox(layout.NewSpacer(), aboutButton, layout.NewSpacer())
	// timerContainer := container.NewHBox(layout.NewSpacer(), starTimerButton, stopTimerButton, resetTimerButton, hideTimerButton, layout.NewSpacer())
	// container := container.NewVBox(scoreboardLabel, blueContainer, goldContainer, scoreboardContainer, resetButtonContainer, link, timerContainer, aboutContainer)
	container := container.NewVBox(scoreboardLabel, blueContainer, goldContainer, scoreboardContainer, resetButtonContainer, link, aboutContainer)
	return container
}

func StartScoreboard(w fyne.Window) {
	w.Resize(fyne.NewSize(400, 500))
	ScoreboardUIContainer := ScoreboardContent(w, AddScoreboardHotkeys)
	w.SetContent(ScoreboardUIContainer)
}

func ProgressIndicator() *fyne.Container {
	label := widget.NewLabel("Fetching data...")
	label.Alignment = fyne.TextAlignCenter
	label.TextStyle.Bold = true
	infinite := widget.NewProgressBarInfinite()
	container := container.NewVBox(label, infinite)
	return container
}

func AboutPage() *fyne.Container {
	label := widget.NewLabel("About KQB Scoreboard")
	label.Alignment = fyne.TextAlignCenter
	label.TextStyle.Bold = true

	author := widget.NewLabel("Author: Prosive")
	author.Alignment = fyne.TextAlignLeading
	githubURL, _ := url.Parse("https://github.com/achhabra2/kqb-scoreboard")
	githubWidget := widget.NewHyperlink("Github Repo", githubURL)
	hotkeys := widget.NewLabel("Hotkeys: \n CTRL + Shift + R: Reset Timer \n CTRL + Shift + T: Toggle Timer Start / Stop \n CTRL + Shift + B: Increment Blue Wins \n CTRL + Shift + G: Increment Gold Wins")
	hotkeys.Alignment = fyne.TextAlignLeading
	container := container.NewVBox(label, author, githubWidget, hotkeys)
	return container
}

func ShowAboutWindow() {
	aboutWindow := FyneApp.NewWindow("About")
	aboutWindow.SetContent(AboutPage())
	aboutWindow.CenterOnScreen()
	aboutWindow.Show()
}

func CheckForUpdates(w fyne.Window) {
	shouldUpdate, latestVersion := checkForUpdate()
	if shouldUpdate {
		updateMessage := fmt.Sprintf("New Version Available, would you like to update to v%s", latestVersion)
		confirmDialog := dialog.NewConfirm("Update Checker", updateMessage, func(action bool) {
			if action {
				log.Println("Update clicked")
				updated := doSelfUpdate()
				if updated {
					updatedDialog := dialog.NewInformation("Update Status", "Update Succeeded, please restart", w)
					updatedDialog.Show()
				} else {
					updatedDialog := dialog.NewInformation("Update Status", "Update Failed", w)
					updatedDialog.Show()
				}
			}
		}, w)
		confirmDialog.Show()
	}
}

func formatTeamName(name string) string {
	return truncate.Truncate(name, 18, "...", truncate.PositionEnd)
}
