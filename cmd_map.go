package main

import (
	"fmt"
)

func mapCommand(cfg *config) error {
	locationAreas, err := cfg.pokeapiClient.GetLocationAreas(cfg.nextLocAreasURL)
	if err != nil {
		return err
	}

	cfg.nextLocAreasURL = locationAreas.Next
	cfg.prevLocAreasURL = locationAreas.Previous

	fmt.Println("")
	for _, area := range locationAreas.Results {
		fmt.Println(area.Name)
	}
	fmt.Println("")

	return nil
}
