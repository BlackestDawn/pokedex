package main

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Prints help text",
			callback:    commandHelp,
		},
		"clear": {
			name:        "clear",
			description: "Clears screen/terminal of text",
			callback:    commandClear,
		},
		"map": {
			name:        "map",
			description: "Get next page of locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get previous page of locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Shows pokemon in an area",
			callback:    commandExplore,
		},
		"exit": {
			name:        "exit",
			description: "Exits the program",
			callback:    commandExit,
		},
	}
}