package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/dothedada/pokemoncli/internals/pokeapi"
	"github.com/dothedada/pokemoncli/internals/pokecache"
)

type config struct {
	prevURL *string
	nextURL *string
	client  pokeapi.Client
	cache   pokecache.Cache
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
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

		if command, ok := getCommand()[commandName]; ok {
			err := command.callback(conf)
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
		"map": {
			name:        "map",
			description: "Shows the locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "map",
			description: "Shows the previous locations",
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
