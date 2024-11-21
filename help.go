package main

import "fmt"

// displayHelp informs the user about our hardcoded functions
func commandHelp() {
	fmt.Println("Welcome to Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Println("help:    - Show available commands")
	fmt.Println("map:     - Print the names of the next 20 location areas")
	fmt.Println("mapb:    - Print the names of the previous 20 location areas")
	fmt.Println("clear:   - Clear the terminal screen")
	fmt.Println("exit:    - Exits the program")
}
