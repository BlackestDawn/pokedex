package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedex/internal/pokeapi"
	"strings"
	"time"
)

// cleanInput preprocesses input to the db repl
func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func main() {
	// Some vars
	clientAPI := pokeapi.NewAPIClient(5 * time.Second)
	cfg := &config{
		pokeApiClient: clientAPI,
	}

	// Begin the repl loop
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\n", cliName, "> ")
		reader.Scan()
		words := cleanInput(reader.Text())
		cmdName := words[0]
		cmd, exists := getCommands()[cmdName]
		if exists {
			err := cmd.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("command not found :", cmd)
		}
	}
}
