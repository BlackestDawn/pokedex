package main

import "pokedex/internal/pokeapi"

type config struct {
	pokeApiClient pokeapi.APIClient
	nextLocURL    *string
	prevLocURL    *string
	params        []string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}
