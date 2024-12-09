package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/migomi3/pokedex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world. Each subsequent call displays the next 20 locations, and so on",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Map back; Displays the previous 20 locations",
			callback:    commandMapB,
		},
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

func commandHelp(cfg *Config) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")

	for _, val := range getCommands() {
		fmt.Printf("\t%s: %s\n", val.name, val.description)
	}

	return nil
}

func commandExit(cfg *Config) error {
	defer os.Exit(0)
	return nil
}

func commandMap(cfg *Config) error {
	locations, err := pokeapi.GetAreas(*cfg.nextLocationURL)
	if err != nil {
		return err
	}

	cfg.nextLocationURL = locations.Next
	cfg.prevLocationURL = locations.Previous

	for _, result := range locations.Results {
		fmt.Println(result.Name)
	}

	return nil
}

func commandMapB(cfg *Config) error {
	if cfg.prevLocationURL == nil {
		return errors.New("already on the first page")
	}

	locations, err := pokeapi.GetAreas(*cfg.prevLocationURL)
	if err != nil {
		return err
	}

	cfg.nextLocationURL = locations.Next
	cfg.prevLocationURL = locations.Previous

	for _, result := range locations.Results {
		fmt.Println(result.Name)
	}

	return nil
}
