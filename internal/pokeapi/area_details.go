package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *APIClient) GetAreaDetails(area string) (LocationAreaDetails, error) {
	url := baseURL + "location-area/" + area

	var data []byte
	var err error
	data, exists := c.apiCache.Get(url)
	if !exists {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return LocationAreaDetails{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return LocationAreaDetails{}, err
		}
		defer resp.Body.Close()

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return LocationAreaDetails{}, err
		}

		c.apiCache.Add(url, data)
	}

	ret_val := LocationAreaDetails{}
	err = json.Unmarshal(data, &ret_val)
	if err != nil {
		return LocationAreaDetails{}, err
	}

	return ret_val, nil
}
