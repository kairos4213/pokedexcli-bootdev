package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationAreas(url *string) (LocationAreas, error) {
	locAreasEndPnt := apiURL + "location-area/"

	if url != nil {
		locAreasEndPnt = *url
	}

	cacheData, ok := c.cache.Get(locAreasEndPnt)
	if ok {
		fmt.Println("cache result")
		var locationAreas LocationAreas
		if err := json.Unmarshal(cacheData, &locationAreas); err != nil {
			return LocationAreas{}, fmt.Errorf("error unmarshaling cache: %w", err)
		}
		return locationAreas, nil
	}
	fmt.Println("api result")

	req, err := http.NewRequest("GET", locAreasEndPnt, nil)
	if err != nil {
		return LocationAreas{}, fmt.Errorf("error creating request: %w", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreas{}, fmt.Errorf("error sending request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return LocationAreas{}, fmt.Errorf("status code: %v", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreas{}, fmt.Errorf("error reading response: %w", err)
	}

	var locationAreas LocationAreas
	if err := json.Unmarshal(data, &locationAreas); err != nil {
		return LocationAreas{}, fmt.Errorf("error unmarshaling data: %w", err)
	}

	c.cache.Add(locAreasEndPnt, data)

	return locationAreas, nil
}
