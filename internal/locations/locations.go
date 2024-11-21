package locations

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetBatchAreas(url string, cfg *LocationPaginate) ([]string, error) {
	res, err := http.Get(url)
	if err != nil {
		return []string{}, fmt.Errorf("error fetching data: %w", err)
	}
	if res.StatusCode < 200 || res.StatusCode > 299 {
		return []string{}, fmt.Errorf("error fetching data: %v", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return []string{}, fmt.Errorf("error reading data: %w", err)
	}
	defer res.Body.Close()

	var data LocationAreaBatch
	err = json.Unmarshal(body, &data)
	if err != nil {
		return []string{}, fmt.Errorf("error preparing data: %w", err)
	}

	ret_val := make([]string, 0)

	for _, place := range data.Results {
		ret_val = append(ret_val, place.Name)
	}
	cfg.NextPage = data.Next
	cfg.PrevPage = data.Previous

	return ret_val, nil
}
