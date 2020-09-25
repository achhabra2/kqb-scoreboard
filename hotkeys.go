package main

import (
	"log"

	hook "github.com/robotn/gohook"
)

func AddEventHotkeys() {
	mode := 0

	hook.Register(hook.KeyDown, []string{"q", "ctrl", "shift"}, func(e hook.Event) {
		log.Println("ctrl-shift-q")
		hook.End()
	})

	hook.Register(hook.KeyDown, []string{"r", "ctrl", "shift"}, func(e hook.Event) {
		log.Println("Reset Timer Hotkey Pressed")
		UpdateTimer("ResetTimer")
	})

	hook.Register(hook.KeyDown, []string{"t", "ctrl", "shift"}, func(e hook.Event) {
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

	s := hook.Start()
	// defer robotgo.EventEnd()
	<-hook.Process(s)
}

func AddScoreboardHotkeys(UIUpdate func()) {
	hook.Register(hook.KeyDown, []string{"b", "ctrl", "shift"}, func(e hook.Event) {
		log.Println("Increment Blue Hotkey Presssed")
		s.IncrementHome()
		UpdateScoreBoard(&s)
		UIUpdate()
	})

	hook.Register(hook.KeyDown, []string{"g", "ctrl", "shift"}, func(e hook.Event) {
		log.Println("Increment Gold Hotkey Pressed")
		s.IncrementAway()
		UpdateScoreBoard(&s)
		UIUpdate()
	})
}
