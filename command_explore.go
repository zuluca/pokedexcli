package main

import (
	"encoding/json"
	"fmt"

	"github.com/zuluca/pokedexcli/internal/pokeapi"
)

func commandExplore(cfg *config, args []string) error {
	if len(args) < 1 {
		fmt.Println("Usage: explore <location-area>")
		return nil
	}

	location := args[0]
	fmt.Printf("Exploring %s...\n", location)

	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", location)
	body, err := cfg.client.FetchData(url)
	if err != nil {
		return err
	}

	var area pokeapi.LocationArea
	if err := json.Unmarshal(body, &area); err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, encounter := range area.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}
