package main

import (
	"fmt"

	"github.com/kairos4213/pokedexcli-bootdev/internal"
)

func mapCommand(c *config) error {
	locationAreas, err := internal.GetLocationAreas(c.next)
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
