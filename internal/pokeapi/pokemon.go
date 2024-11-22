package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *APIClient) GetPokemonDetails(pokemon string) (PokemonDetails, error) {
	url := pokemonURL + pokemon

	var data []byte
	var err error
	data, exists := c.apiCache.Get(url)
	if !exists {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return PokemonDetails{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return PokemonDetails{}, err
		}
		defer resp.Body.Close()

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return PokemonDetails{}, err
		}

		c.apiCache.Add(url, data)
	}

	ret_val := PokemonDetails{}
	err = json.Unmarshal(data, &ret_val)
	if err != nil {
		return PokemonDetails{}, err
	}

	return ret_val, nil
}

func (c *APIClient) AddPokemon(pokemon PokemonDetails) {
	c.pokemons[pokemon.Name] = pokemon
}
