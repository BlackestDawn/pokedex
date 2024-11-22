package main

import (
	"fmt"
	"math/rand/v2"
)

func commandCatch(cfg *config) error {
	if len(cfg.params) == 0 {
		return fmt.Errorf("no pokemon given to catch")
	}
	resp, err := cfg.pokeApiClient.GetPokemonDetails(cfg.params[0])
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a pokeball at %s...\n", resp.Name)
	catchChance := rand.IntN(resp.BaseExperience)

	if catchChance <= catchPoint {
		cfg.pokeApiClient.AddPokemon(resp)
		fmt.Println(resp.Name, "was caught!")
	} else {
		fmt.Println(resp.Name, "escaped!")
	}

	return nil
}
