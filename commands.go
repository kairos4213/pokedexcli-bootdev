package main

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
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
		"map": {
			name:        "map",
			description: "Displays the next 20 locations.",
			callback:    mapCommand,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations. It's a way to go back.",
			callback:    mapbCommand,
		},
	}

	return commands
}
