package main

import (
	"fmt"
	"strings"
)

func commandInspect(cfg *config, args []string) error {
	if len(args) == 0 {
		fmt.Println("Please specify a Pokemon to inspect")
		return nil
	}

	name := strings.ToLower(args[0])

	pokemon, ok := cfg.Pokedex[name]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)

	fmt.Println("Stats:")
	for statName, statValue := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", statName, statValue)
	}

	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t)
	}

	return nil
}
