package main

import (
	"time"

	"github.com/saga-sanga/pokedexcli/internal/pokeapi"
	"github.com/saga-sanga/pokedexcli/internal/pokecache"
)

func main() {
	pokecache := pokecache.NewCache(5 * time.Second)
	pokeClient := pokeapi.NewClient(5*time.Second, pokecache)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
