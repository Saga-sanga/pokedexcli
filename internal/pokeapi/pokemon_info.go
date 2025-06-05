package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) PokemonInfo(name string) (RespPokemonInfo, error) {
	url := baseURL + "/pokemon/" + name

	pokemonResp := RespPokemonInfo{}
	pokemonData, exist := c.pokecache.Get(url)
	if exist {
		err := json.Unmarshal(pokemonData, &pokemonResp)
		if err != nil {
			return RespPokemonInfo{}, err
		}
		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokemonInfo{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokemonInfo{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespPokemonInfo{}, err
	}

	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return RespPokemonInfo{}, err
	}

	c.pokecache.Add(url, dat)

	return pokemonResp, nil
}
