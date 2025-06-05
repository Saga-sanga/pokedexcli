package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, pokemon string) error {
	pokeInfo, err := cfg.pokeapiClient.PokemonInfo(pokemon)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %v...\n", pokeInfo.Name)
	chance := rand.Intn(pokeInfo.BaseExperience / 20)
	if chance == 1 {
		// Success
		fmt.Printf("%v was caught!\n", pokeInfo.Name)
		cfg.pokedex[pokeInfo.Name] = pokeInfo
	} else {
		fmt.Printf("%v escaped!\n", pokeInfo.Name)
	}
	return nil
}
