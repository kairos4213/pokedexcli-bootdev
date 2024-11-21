package main

import (
	"fmt"
)

func helpCommand() error {
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Println("Usage")
	fmt.Println("")

	commands := getCommandsMap()
	for command := range commands {
		name := commands[command].name
		description := commands[command].description
		fmt.Printf("%s: %s", name, description)
		fmt.Println("")
	}
	fmt.Println("")

	return nil
}
