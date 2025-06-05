package main

import "fmt"

func commandExplore(cfg *config, location string) error {
	resp, err := cfg.pokeapiClient.ListPokemon(location)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %v...\n", location)
	fmt.Println("Found Pokemon:")
	for _, encounter := range resp.PokemonEncounters {
		fmt.Printf(" - %v\n", encounter.Pokemon.Name)
	}

	return nil
}
