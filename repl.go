package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/zuluca/pokedexcli/internal/pokeapi"
)

func startREPL(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		line := reader.Text()

		words := cleanInput(line)
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		cmd, ok := getCommands()[commandName]

		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		if err := cmd.callback(cfg, words[1:]); err != nil {
			fmt.Printf("Error: %v\n", err)
		}

	}
}

func cleanInput(text string) []string {
	// Trim leading/trailing spaces
	trimmed := strings.TrimSpace(text)

	// Lowercase everything
	lowered := strings.ToLower(trimmed)

	// Split on whitespace
	words := strings.Fields(lowered)

	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

type locationAreaResp struct {
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type config struct {
	next     *string
	previous *string
	client   *pokeapi.Client
	Pokedex  map[string]Pokemon
}

type Pokemon struct {
	Name           string
	Height         int
	Weight         int
	Stats          map[string]int
	Types          []string
	BaseExperience int
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Go back to the previous 20 location areas",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore",
			description: "Explore a location area and list Pokémon",
			callback: func(cfg *config, args []string) error {
				// pass along the rest of the words (args)
				return commandExplore(cfg, args)
			},
		},
		"catch": {
			name:        "catch",
			description: "Catch a Pokémon and add it to your Pokedex",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a caught Pokémon in your Pokedex",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all caught Pokémon in your Pokedex",
			callback:    commandPokedex,
		},
	}
}
