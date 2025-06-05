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
	pokedex          map[string]pokeapi.RespPokemonInfo
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

		var param string
		commandWord := words[0]
		if len(words) > 1 {
			param = words[1]
		}
		command, exists := getCommands()[commandWord]
		if exists {
			err := command.callback(cfg, param)
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
	callback    func(*config, string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			description: "Displays the help message",
			name:        "help",
			callback:    commandHelp,
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
		"explore": {
			description: "Get a list of all pokemon located in an area",
			name:        "explore",
			callback:    commandExplore,
		},
		"catch": {
			description: "Catch a given pokemon",
			name:        "catch",
			callback:    commandCatch,
		},
	}
}
