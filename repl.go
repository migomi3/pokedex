package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/migomi3/pokedex/internal/pokecache"
)

func startRepl() error {
	scanner := bufio.NewScanner(os.Stdin)
	startURL := "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"

	cfg := &Config{
		cache:           pokecache.NewCache(5 * time.Minute),
		nextLocationURL: &startURL,
		prevLocationURL: nil,
	}

	fmt.Println()

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		input := scanner.Text()
		cleanedCommand := cleanInput(input)

		command, exists := getCommands()[cleanedCommand]
		if exists {
			err := command.callback(cfg)

			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown Command")
		}

	}
}

func cleanInput(s string) string {
	lowered := strings.ToLower(s)
	cleanedText := strings.Fields(lowered)
	return cleanedText[0]
}
