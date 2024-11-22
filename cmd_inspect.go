package main

import "fmt"

func commandInspect(cfg *config) error {
	name := cfg.params[0]

	pokemon, exists := cfg.pokeApiClient.GetPokemon(name)
	if !exists {
		return fmt.Errorf("no pokemon caught by that name")
	}

	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  - %s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, stat := range pokemon.Types {
		fmt.Printf("  - %s\n", stat.Type.Name)
	}

	return nil
}
