package main

import (
	"fmt"
)

func helpCommand(c *config) error {
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Println("Usage")
	fmt.Println("")

	commands := getCommandsMap()
	for id := range commands {
		name := commands[id].name
		description := commands[id].description
		fmt.Printf("%s: %s", name, description)
		fmt.Println("")
	}
	fmt.Println("")

	return nil
}
