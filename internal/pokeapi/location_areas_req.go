package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) GetLocationAreas(url *string) (LocationAreas, error) {
	locAreasEndPnt := apiURL + "location-area/"

	if url != nil {
		locAreasEndPnt = *url
	}

	req, err := http.NewRequest("GET", locAreasEndPnt, nil)
	if err != nil {
		return LocationAreas{}, fmt.Errorf("error creating request: %w", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreas{}, fmt.Errorf("error sending request: %w", err)
	}

	if res.StatusCode > 399 {
		return LocationAreas{}, fmt.Errorf("status code: %v", res.StatusCode)
	}
	defer res.Body.Close()

	var locationAreas LocationAreas
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&locationAreas); err != nil {
		return LocationAreas{}, fmt.Errorf("error decoding response: %w", err)
	}

	return locationAreas, nil
}
