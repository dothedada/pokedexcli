package main

import (
	"github.com/dothedada/pokemoncli/internals/pokeapi"
	"github.com/dothedada/pokemoncli/internals/pokecache"
)

func main() {
	client := pokeapi.NewClient(pokeapi.Timeout, pokecache.CacheTime)
	pokedex := map[string]pokeapi.PokemonData{}

	conf := &config{
		client:  client,
		pokedex: pokedex,
	}

	repl(conf)
}
