package main

import "fmt"

// displayHelp informs the user about our hardcoded functions
func commandHelp(cfg *config) error {
	fmt.Println("Welcome to Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, val := range getCommands() {
		fmt.Printf("%-10s- %s\n", val.name, val.description)
	}
	return nil
}
