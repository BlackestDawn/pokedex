package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedex/internal/locations"
	"strings"
)

// printPrompt displays the repl prompt at the start of each loop
func printPrompt() {
	fmt.Print("\n", cliName, "> ")
}

// printUnkown informs the user about invalid commands
func printUnknown(text string) {
	fmt.Println(text, ": command not found")
	commandHelp()
}

// cleanInput preprocesses input to the db repl
func cleanInput(text string) string {
	output := strings.TrimSpace(text)
	output = strings.ToLower(output)
	return output
}

// parseCmd parses command given
func parseCmd(cmd string) {
	printUnknown(cmd)
}

func main() {
	// Some vars
	tmpStr := new(string)
	*tmpStr = locationAreaAPI
	areaLocs := &locations.LocationPaginate{
		NextPage: tmpStr,
		PrevPage: nil,
	}

	// Begin the repl loop
	reader := bufio.NewScanner(os.Stdin)
	printPrompt()
	for reader.Scan() {
		text := cleanInput(reader.Text())
		switch text {
		case "exit":
			commandExit()
		case "help":
			commandHelp()
		case "clear":
			commandClear()
		case "map":
			commandMap(areaLocs, true)
		case "mapb":
			commandMap(areaLocs, false)
		default:
			// Pass the command to the parser
			parseCmd(text)
		}
		printPrompt()
	}
	// Print an additional line if we encountered an EOF character
	fmt.Println()
}
