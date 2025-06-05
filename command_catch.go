package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	pokeInfo, err := cfg.pokeapiClient.PokemonInfo(name)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %v...\n", pokeInfo.Name)
	chance := rand.Intn(pokeInfo.BaseExperience / 20)
	if chance == 1 {
		// Success
		fmt.Printf("%v was caught!\n", pokeInfo.Name)
		cfg.caughtPokemon[pokeInfo.Name] = pokeInfo
	} else {
		fmt.Printf("%v escaped!\n", pokeInfo.Name)
	}
	return nil
}
