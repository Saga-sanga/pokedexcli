package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListPokemon -
func (c *Client) ListPokemon(location string) (RespPokeLocations, error) {
	url := baseURL + "/location-area/" + location

	pokemonResp := RespPokeLocations{}

	// Check cache
	// check if url already exists
	data, exists := c.pokecache.Get(url)

	// return data without fetching if exists
	if exists {
		err := json.Unmarshal(data, &pokemonResp)
		if err != nil {
			return RespPokeLocations{}, nil
		}
		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokeLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokeLocations{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespPokeLocations{}, err
	}

	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return RespPokeLocations{}, err
	}

	c.pokecache.Add(url, dat)

	return pokemonResp, nil
}
