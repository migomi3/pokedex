package main

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"map": {
			name:        "map",
			description: "Displays names of 20 location areas in Pokemon world. Each subsequent call displays the next 20, and so on",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Map back; Displays the previous 20 locations",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Returns names of pokemon found in location area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Try and take a chance to catch a pokemon from the arg",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "see details about a Pokemon you've caught",
			callback:    commandInspect,
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
