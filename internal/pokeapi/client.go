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
}

// Create new client
func NewAPIClient(timeout time.Duration) APIClient {
	return APIClient{
		httpClient: http.Client{
			Timeout: timeout,
		},
		apiCache: pokecache.NewCache(cacheTimeout),
	}
}
