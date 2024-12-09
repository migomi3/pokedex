package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() error {
	scanner := bufio.NewScanner(os.Stdin)
	startURL := "https://pokeapi.co/api/v2/location-area/"

	cfg := &Config{
		nextLocationURL: &startURL,
		prevLocationURL: nil,
	}

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
