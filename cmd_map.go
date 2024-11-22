package main

import (
	"fmt"
)

// commandMap fetch next 20 location areas
func commandMap(cfg *config) error {
	resp, err := cfg.pokeApiClient.GetLocationAreas(cfg.nextLocURL)
	if err != nil {
		return err
	}

	cfg.nextLocURL = resp.Next
	cfg.prevLocURL = resp.Previous

	for _, loc := range resp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

// commandMap fetch next 20 location areas
func commandMapb(cfg *config) error {
	if cfg.prevLocURL == nil {
		return fmt.Errorf("already on first page")
	}

	resp, err := cfg.pokeApiClient.GetLocationAreas(cfg.prevLocURL)
	if err != nil {
		return err
	}

	cfg.nextLocURL = resp.Next
	cfg.prevLocURL = resp.Previous

	for _, loc := range resp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
