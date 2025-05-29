package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Location struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

func displayMap(out io.Writer) error {

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	req, err := http.NewRequest("GET", "https://pokeapi.co/api/v2/location", nil)
	if err != nil {
		return err
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	var location Location

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&location); err != nil {
		return err
	}

	for _, v := range location.Results {
		fmt.Fprintf(out, "%v\n", v.Name)
	}

	return nil
}
