package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/kairos4213/pokedexcli-bootdev/internal/pokeapi"
)

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex > ")

		scanner.Scan()
		userInput := scanner.Text()
		userInput = strings.TrimSpace(userInput)
		if len(userInput) == 0 {
			continue
		}

		commands := getCommandsMap()
		command, ok := commands[userInput]
		if !ok {
			fmt.Println("invalid command")
			continue
		}

		err := command.callback(&cfg)
		if err != nil {
			fmt.Println(err)
		}
	}
}
