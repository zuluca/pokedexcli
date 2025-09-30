package main

import (
	"time"

	"github.com/zuluca/pokedexcli/internal/pokeapi"
)

func main() {
	cfg := &config{
		client:  pokeapi.NewClient(5 * time.Second), // keep your client
		Pokedex: make(map[string]Pokemon),           // add Pokedex map
	}

	startREPL(cfg)
}
