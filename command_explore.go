package main

import "fmt"

func commandExplore(cfg *config, location string) error {
	resp, err := cfg.pokeapiClient.ListPokemon(location)
	if err != nil {
		return err
	}

	for _, encounter := range resp.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}

	return nil
}
