package main

import (
	"errors"
	"fmt"
)

func mapbCommand(cfg *config) error {
	if cfg.prevLocAreasURL == nil {
		return errors.New("cannot go back from first page")
	}

	locationAreas, err := cfg.pokeapiClient.GetLocationAreas(cfg.prevLocAreasURL)
	if err != nil {
		return err
	}

	cfg.nextLocAreasURL = locationAreas.Next
	cfg.prevLocAreasURL = locationAreas.Previous

	fmt.Println("")
	for _, result := range locationAreas.Results {
		fmt.Println(result.Name)
	}
	fmt.Println("")

	return nil
}
