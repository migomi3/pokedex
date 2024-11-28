package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() error {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		input := scanner.Text()
		cleanedCommand := cleanInput(input)

		command, exists := getCommands()[cleanedCommand]
		if exists {
			err := command.callback()

			if err != nil {
				return err
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
