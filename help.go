package main

import "fmt"

// displayHelp informs the user about our hardcoded functions
func commandHelp() error {
	fmt.Println("Welcome to Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Println("help:    - Show available commands")
	fmt.Println("clear:   - Clear the terminal screen")
	fmt.Println("exit:    - Exits the program")

	return nil
}
