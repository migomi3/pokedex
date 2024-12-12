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
	fmt.Println("Closing the Pokedex... Goodbye!")
	defer os.Exit(0)
	return nil
}

func commandMap(cfg *Config) error {
	var locations pokeapi.LocationAreaRes
	var err error

	if body, ok := cfg.cache.Get(*cfg.nextLocationURL); ok {
		locations, err = pokeapi.UnmarshalAreas(body)
		if err != nil {
			return err
		}
	} else {
		locations, err = pokeapi.GetAreas(*cfg.nextLocationURL, &cfg.cache)
		if err != nil {
			return err
		}
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

	var locations pokeapi.LocationAreaRes
	var err error

	if body, ok := cfg.cache.Get(*cfg.prevLocationURL); ok {
		locations, err = pokeapi.UnmarshalAreas(body)
		if err != nil {
			return err
		}
	} else {
		locations, err = pokeapi.GetAreas(*cfg.prevLocationURL, &cfg.cache)
		if err != nil {
			return err
		}
	}

	cfg.nextLocationURL = locations.Next
	cfg.prevLocationURL = locations.Previous

	for _, result := range locations.Results {
		fmt.Println(result.Name)
	}

	return nil
}
