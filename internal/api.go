package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LocationAreas struct {
	Count    int     `json:"id"`
	Next     string  `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocationAreas(url string) (LocationAreas, error) {
	res, err := http.Get(url)
	if err != nil {
		return LocationAreas{}, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	var locationAreas LocationAreas
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&locationAreas); err != nil {
		return LocationAreas{}, fmt.Errorf("error decoding response: %w", err)
	}

	return locationAreas, nil
}
