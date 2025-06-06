package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/saga-sanga/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	caughtPokemon    map[string]pokeapi.Pokemon
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandWord := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := getCommands()[commandWord]
		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}

	}
}

func cleanInput(text string) []string {
	splitWords := strings.Fields(strings.ToLower(text))
	return splitWords
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			description: "Displays the help message",
			name:        "help",
			callback:    commandHelp,
		},
		"explore": {
			description: "explore <location_name>",
			name:        "explore",
			callback:    commandExplore,
		},
		"catch": {
			description: "Catch a given pokemon",
			name:        "catch <pokemon_name>",
			callback:    commandCatch,
		},
		"inspect": {
			description: "Inspect a caught pokemon",
			name:        "inspect <pokemon_name>",
			callback:    commandInspect,
		},
		"map": {
			description: "Get the next page of locations",
			name:        "map",
			callback:    commandMapf,
		},
		"mapb": {
			description: "Get the previous page of locations",
			name:        "mapb",
			callback:    commandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
