package main

import "fmt"

func CommandList(cfg *config) error {
	pokeList := cfg.pokeApiClient.ListPokemons()
	if len(pokeList) == 0 {
		return fmt.Errorf("no entries in Pokedex")
	}
	fmt.Println("Your Pokedex:")
	for _, name := range pokeList {
		fmt.Println("  -", name)
	}
	return nil
}
