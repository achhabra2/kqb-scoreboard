package main

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func PromptIGLCustom() {
	// Each template displays the data received from the prompt with some formatting.
	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
		Confirm: "Is this an IGL Match or Custom Match?",
	}

	prompt := promptui.Prompt{
		Label:     "Select Match Type?",
		Templates: templates,
		IsConfirm: true,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	// The result of the prompt, if valid, is displayed in a formatted message.
	fmt.Printf("You answered %s\n", result)
	return
}
