package pokeapi

import "time"

// URLs for handling pokeAPI lookups
const baseURL string = "https://pokeapi.co/api/v2/"
const locationURL string = baseURL + "location-area/"
const pokemonURL string = baseURL + "pokemon/"

// cacheTimeout is the time before cached items get removed
const cacheTimeout = 30 * time.Second
