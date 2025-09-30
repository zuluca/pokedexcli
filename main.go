package main

import (
	"time"

	"github.com/zuluca/pokedexcli/internal/pokeapi"
)

func main() {
	cacheInterval := 5 * time.Second
	client := pokeapi.NewClient(cacheInterval)

	cfg := &config{
		client: client,
	}

	startREPL(cfg)
}
