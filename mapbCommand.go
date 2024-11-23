package main

import (
	"errors"
	"fmt"

	"github.com/kairos4213/pokedexcli-bootdev/internal"
)

func mapbCommand(c *config) error {
	if c.previous == "" {
		return errors.New("cannot return past 0")
	}

	locationAreas, err := internal.GetLocationAreas(c.previous)
	if err != nil {
		return err
	}

	c.next = locationAreas.Next
	c.previous = locationAreas.Previous

	fmt.Println("")
	for _, result := range locationAreas.Results {
		fmt.Println(result.Name)
	}
	fmt.Println("")

	return nil
}
