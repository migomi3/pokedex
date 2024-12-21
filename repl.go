package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/migomi3/pokedex/internal/pokeapi"
	"github.com/migomi3/pokedex/internal/pokecache"
)

func startRepl() error {
	scanner := bufio.NewScanner(os.Stdin)
	baseURL := "https://pokeapi.co/api/v2"
	pokedex := make(map[string]pokeapi.Pokemon)

	cfg := &Config{
		cache:           pokecache.NewCache(5 * time.Minute),
		baseURL:         &baseURL,
		nextLocationURL: nil,
		prevLocationURL: nil,
		pokedex:         pokedex,
	}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		input := scanner.Text()
		cleanedCommand := cleanInput(input)
		var arg string
		if len(cleanedCommand) == 0 {
			continue
		}

		if len(cleanedCommand) > 1 {
			arg = cleanedCommand[1]
		}

		command, exists := getCommands()[cleanedCommand[0]]
		if exists {
			err := command.callback(cfg, arg)

			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown Command")
		}

	}
}

func cleanInput(s string) []string {
	lowered := strings.ToLower(s)
	cleanedText := strings.Fields(lowered)
	return cleanedText
}
