package main

import (
	"errors"
	"fmt"
	"math/rand"
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

	fmt.Printf("Exploring %s...\n", location.Name)
	fmt.Println("Found Pokemon:")

	for _, pokemonEncounter := range location.PokemonEncounters {
		fmt.Printf("\t- %s\n", pokemonEncounter.Pokemon.Name)
	}

	return nil
}

func commandCatch(cfg *Config, pokeId string) error {
	if pokeId == "" {
		return errors.New("Need a pokemon Name/id as an argument to catch one")
	}

	url := *cfg.baseURL + "/pokemon/" + pokeId

	var pokemon pokeapi.Pokemon
	var err error

	if body, ok := cfg.cache.Get(url); ok {
		pokemon, err = pokeapi.UnmarshalPokemon(body)
		if err != nil {
			return err
		}
	} else {
		pokemon, err = pokeapi.GetPokemon(&url, &cfg.cache)
		if err != nil {
			return err
		}
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	odds := rand.Intn(750)
	if pokemon.BaseExperience < odds {
		fmt.Printf("%s successfully caught!\n", pokemon.Name)
		cfg.pokedex[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}

func commandInspect(cfg *Config, pokeName string) error {
	if pokeName == "" {
		return errors.New("Need the name of a pokemon you've caught")
	}

	pokemon, exists := cfg.pokedex[pokeName]
	if !exists {
		return fmt.Errorf("%s not found in pokedex", pokeName)
	}

	fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\nStats:\n", pokemon.Name, pokemon.Height, pokemon.Weight)
	for _, stat := range pokemon.Stats {
		fmt.Printf("\t-%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")

	for _, pokeType := range pokemon.Types {
		fmt.Printf("\t-%s\n", pokeType.Type.Name)
	}

	return nil
}
