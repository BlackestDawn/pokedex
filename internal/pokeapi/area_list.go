package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *APIClient) GetLocationAreas(pageURL *string) (LocationAreaBatch, error) {
	url := baseURL + "location-area/"
	if pageURL != nil {
		url = *pageURL
	}

	var data []byte
	var err error
	data, exists := c.apiCache.Get(url)
	if !exists {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return LocationAreaBatch{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return LocationAreaBatch{}, err
		}
		defer resp.Body.Close()

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return LocationAreaBatch{}, err
		}

		c.apiCache.Add(url, data)
	}

	ret_val := LocationAreaBatch{}
	err = json.Unmarshal(data, &ret_val)
	if err != nil {
		return LocationAreaBatch{}, err
	}

	return ret_val, nil
}
