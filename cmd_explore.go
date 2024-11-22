package main

import "fmt"

func commandExplore(cfg *config) error {
	if len(cfg.params[0]) == 0 {
		return fmt.Errorf("no area given")
	}
	resp, err := cfg.pokeApiClient.GetAreaDetails(cfg.params[0])
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", resp.Name)
	if len(resp.PokemonEncounters) > 0 {
		fmt.Println("Found Pokemon:")
		for _, val := range resp.PokemonEncounters {
			fmt.Println("  -", val.Pokemon.Name)
		}
	} else {
		fmt.Println("No pokemon found")
	}
	return nil
}
