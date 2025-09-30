package main

import (
	"encoding/json"
	"fmt"
)

func commandMap(cfg *config, args []string) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if cfg.next != nil {
		url = *cfg.next
	}

	body, err := cfg.client.FetchData(url)
	if err != nil {
		return err
	}

	var locations locationAreaResp
	if err := json.Unmarshal(body, &locations); err != nil {
		return err
	}

	// Print all names
	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}

	// Update config for pagination
	cfg.next = locations.Next
	cfg.previous = locations.Previous

	return nil
}

func commandMapBack(cfg *config, args []string) error {
	if cfg.previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	body, err := cfg.client.FetchData(*cfg.previous)
	if err != nil {
		return err
	}

	var locations locationAreaResp
	if err := json.Unmarshal(body, &locations); err != nil {
		return err
	}

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}

	cfg.next = locations.Next
	cfg.previous = locations.Previous

	return nil
}
