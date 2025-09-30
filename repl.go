package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/zuluca/pokedexcli/internal/pokeapi"
)

func startREPL() {
	reader := bufio.NewScanner(os.Stdin)
	cfg := &config{} //shared state between commands

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

		if err := cmd.callback(cfg); err != nil {
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
	callback    func(*config) error
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
	}
}
