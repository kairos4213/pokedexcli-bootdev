package main

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommandsMap() map[string]cliCommand {
	commands := map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays information about pokedex",
			callback:    helpCommand,
		},
		"exit": {
			name:        "exit",
			description: "Exit pokedex",
			callback:    exitCommand,
		},
	}

	return commands
}
