package pokeapi

import (
	"net/http"
	"pokedex/internal/pokecache"
	"time"
)

// Base client for the PokeAPI
type APIClient struct {
	httpClient http.Client
	apiCache   *pokecache.PokeCache
	pokemons   map[string]PokemonDetails
}

// Create new client
func NewAPIClient(timeout time.Duration) APIClient {
	return APIClient{
		httpClient: http.Client{
			Timeout: timeout,
		},
		apiCache: pokecache.NewCache(cacheTimeout),
		pokemons: make(map[string]PokemonDetails),
	}
}
