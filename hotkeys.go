package main

import (
	"log"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

func AddEventHotkeys() {
	mode := 0

	robotgo.EventHook(hook.KeyDown, []string{"q", "ctrl", "shift"}, func(e hook.Event) {
		log.Println("ctrl-shift-q")
		robotgo.EventEnd()
	})

	robotgo.EventHook(hook.KeyDown, []string{"r", "ctrl", "shift"}, func(e hook.Event) {
		log.Println("Reset Timer Hotkey Pressed")
		UpdateTimer("ResetTimer")
	})

	robotgo.EventHook(hook.KeyDown, []string{"t", "ctrl", "shift"}, func(e hook.Event) {
		log.Println("Toggle Start/Stop Timer Hotkey Pressed")
		switch mode {
		case 0:
			UpdateTimer("StartTimer")
			mode = 1
			break
		case 1:
			UpdateTimer("StopTimer")
			mode = 0
			break
		default:
			break
		}
	})

	s := robotgo.EventStart()
	// defer robotgo.EventEnd()
	<-robotgo.EventProcess(s)
}

func AddScoreboardHotkeys(UIUpdate func()) {
	robotgo.EventHook(hook.KeyDown, []string{"b", "ctrl", "shift"}, func(e hook.Event) {
		log.Println("Increment Blue Hotkey Presssed")
		s.IncrementHome()
		UpdateScoreBoard(&s)
		UIUpdate()
	})

	robotgo.EventHook(hook.KeyDown, []string{"g", "ctrl", "shift"}, func(e hook.Event) {
		log.Println("Increment Gold Hotkey Pressed")
		s.IncrementAway()
		UpdateScoreBoard(&s)
		UIUpdate()
	})
}
