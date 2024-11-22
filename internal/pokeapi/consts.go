package pokeapi

import "time"

// baseURL is the base URL for the pokeAPI lookups
const baseURL string = "https://pokeapi.co/api/v2/"

// cacheTimeout is the time before cached items get removed
const cacheTimeout = 30 * time.Second
