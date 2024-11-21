package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// cliName is the name used in the repl prompts
var cliName string = "Pokedex"

// printPrompt displays the repl prompt at the start of each loop
func printPrompt() {
	fmt.Print("\n", cliName, "> ")
}

// printUnkown informs the user about invalid commands
func printUnknown(text string) {
	fmt.Println(text, ": command not found")
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
	// Begin the repl loop
	reader := bufio.NewScanner(os.Stdin)
	printPrompt()
	for reader.Scan() {
		text := cleanInput(reader.Text())
		if command, exists := commands[text]; exists {
			// Call a hardcoded function
			command.callback()
		} else {
			// Pass the command to the parser
			parseCmd(text)
		}
		printPrompt()
	}
	// Print an additional line if we encountered an EOF character
	fmt.Println()
}
