package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func startRepl() {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		command, exists := getCommands()[words[0]]
		if exists {
			err := command.callback()
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
			callback: func() error {
				return displayHelp(os.Stdout)
			},
		},
		"map": {
			description: "Displays different locations in the world",
			name:        "map",
			callback: func() error {
				return displayMap(os.Stdout)
			},
		},
	}
}
