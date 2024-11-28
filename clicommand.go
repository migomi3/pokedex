package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")

	for _, val := range getCommands() {
		fmt.Printf("\t%s: %s\n", val.name, val.description)
	}

	return nil
}

func commandExit() error {
	defer os.Exit(0)
	return nil
}
