package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	c := config{next: url, previous: ""}
	commands := getCommandsMap()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex > ")

		scanner.Scan()
		userInput := scanner.Text()

		command := strings.TrimSpace(userInput)
		switch command {
		case "help":
			commands[command].callback(&c)
		case "exit":
			commands[command].callback(&c)
		case "map":
			commands[command].callback(&c)
		case "mapb":
			commands[command].callback(&c)
		default:
			continue
		}
	}
}
