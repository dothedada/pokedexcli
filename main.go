package main

import (
	"github.com/dothedada/pokemoncli/internals/pokeapi"
	"github.com/dothedada/pokemoncli/internals/pokecache"
)

func main() {
	client := pokeapi.NewClient(pokeapi.Timeout, pokecache.CacheTime)
	conf := &config{
		client: client,
	}

	repl(conf)
}
