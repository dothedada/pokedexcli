package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/dothedada/pokemoncli/internals/pokeapi"
)

type config struct {
	prevURL *string
	nextURL *string
	client  pokeapi.Client
	pokedex map[string]pokeapi.PokemonData
}

type cliCommand struct {
	name        string
	description string
	callback    func(conf *config, param ...string) error
}

func repl(conf *config) {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		text := cleanInput(reader.Text())
		if len(text) == 0 {
			continue
		}

		commandName := text[0]
		var args []string

		if len(text) > 1 {
			args = text[1:]
		}

		if command, ok := getCommand()[commandName]; ok {
			err := command.callback(conf, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	lowerTxt := strings.ToLower(text)
	words := strings.Fields(lowerTxt)
	return words
}

func getCommand() map[string]cliCommand {
	return map[string]cliCommand{
		"pokedex": {
			name:        "pokedex",
			description: "Show the list of the Pokemons in the Pokedex",
			callback:    commandPokedex,
		},
		"inspect": {
			name:        "inspect <Pokemon_name>",
			description: "show the data of the specified Pokemon if it is in the Pokedex ",
			callback:    commandInspect,
		},
		"catch": {
			name:        "catch <Pokemon_name>",
			description: "Try to catch the specified Pokemon",
			callback:    commandCatch,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Shows the pokemons in the location name",
			callback:    commandExplore,
		},
		"map": {
			name:        "map",
			description: "Shows the next page of locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Shows the previous page of locations",
			callback:    commandMapBack,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
}
