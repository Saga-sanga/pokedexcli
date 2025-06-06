package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	if val, ok := cfg.caughtPokemon[name]; ok {
		fmt.Printf("Name: %v\n", val.Name)
		fmt.Printf("Height: %v\n", val.Height)
		fmt.Printf("Weight: %v\n", val.Weight)
		fmt.Println("Stats:")
		for _, stat := range val.Stats {
			fmt.Printf("  -%v: %v\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, pokeType := range val.Types {
			fmt.Printf("  - %v\n", pokeType.Type.Name)
		}

	} else {
		fmt.Println("you have not caught that pokemon")
	}

	return nil
}
