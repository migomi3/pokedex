package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/migomi3/pokedex/internal/pokeapi"
)

func commandHelp(cfg *Config, _ string) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")

	for _, val := range getCommands() {
		fmt.Printf("\t%s: %s\n", val.name, val.description)
	}

	return nil
}

func commandExit(cfg *Config, _ string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	defer os.Exit(0)
	return nil
}

func commandMap(cfg *Config, _ string) error {
	var locations pokeapi.LocationAreaRes
	var err error

	url := *cfg.baseURL + "/location-area"
	if cfg.nextLocationURL != nil {
		url = *cfg.nextLocationURL
	}

	if body, ok := cfg.cache.Get(url); ok {
		locations, err = pokeapi.UnmarshalLocationAreasRes(body)
		if err != nil {
			return err
		}
	} else {
		locations, err = pokeapi.GetLocationAreasRes(&url, &cfg.cache)
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

func commandMapB(cfg *Config, _ string) error {
	if cfg.prevLocationURL == nil {
		return errors.New("already on the first page")
	}

	var locations pokeapi.LocationAreaRes
	var err error

	if body, ok := cfg.cache.Get(*cfg.prevLocationURL); ok {
		locations, err = pokeapi.UnmarshalLocationAreasRes(body)
		if err != nil {
			return err
		}
	} else {
		locations, err = pokeapi.GetLocationAreasRes(cfg.prevLocationURL, &cfg.cache)
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

func commandExplore(cfg *Config, area string) error {
	if area == "" {
		return errors.New("explore command needs location-area name or id")
	}

	fmt.Printf("Exploring %s...\n", area)

	url := *cfg.baseURL + "/location-area/" + area

	var location pokeapi.LocationArea
	var err error

	if body, ok := cfg.cache.Get(url); ok {
		location, err = pokeapi.UnmarshalLocationArea(body)
		if err != nil {
			return err
		}
	} else {
		location, err = pokeapi.GetLocationArea(&url, &cfg.cache)
		if err != nil {
			return err
		}
	}

	fmt.Println("Found Pokemon:")

	for _, pokemonEncounter := range location.PokemonEncounters {
		fmt.Printf("\t- %s\n", pokemonEncounter.Pokemon.Name)
	}

	return nil
}
