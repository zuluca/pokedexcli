package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type pokemonResp struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
}

func commandCatch(cfg *config, args []string) error {
	if len(args) < 1 {
		fmt.Println("Usage: catch <pokemon>")
		return nil
	}

	name := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", name)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println("Pokémon not found!")
		return nil
	}

	var p pokemonResp
	if err := json.NewDecoder(resp.Body).Decode(&p); err != nil {
		return err
	}

	// Determine chance of catching (higher base experience = harder)
	rand.Seed(time.Now().UnixNano())
	chance := rand.Intn(100)
	difficulty := p.BaseExperience // higher = harder
	if chance >= difficulty {
		fmt.Printf("%s was caught!\n", p.Name)
		cfg.Pokedex[p.Name] = Pokemon{Name: p.Name, BaseExperience: p.BaseExperience}
	} else {
		fmt.Printf("%s escaped!\n", p.Name)
	}

	return nil
}
