package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	commands := getCommandsMap()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex > ")

		scanner.Scan()
		userInput := scanner.Text()

		command := strings.TrimSpace(userInput)
		switch command {
		case "help":
			commands["help"].callback()
		case "exit":
			commands["exit"].callback()
		default:
			continue
		}
	}
}
