package main

import "fmt"

func commandPokedex(cfg *config, args []string) error {
	if len(cfg.Pokedex) == 0 {
		fmt.Println("Your Pokedex is empty!")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for name := range cfg.Pokedex {
		fmt.Printf(" - %s\n", name)
	}

	return nil
}
