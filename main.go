package main

import (
	"time"

	"github.com/dothedada/pokemoncli/internals/pokeapi"
	"github.com/dothedada/pokemoncli/internals/pokecache"
)

func main() {
	client := pokeapi.NewClient(3 * time.Second)
	cache := pokecache.NewCache(5 * time.Second)
	conf := &config{
		client: client,
		cache:  cache,
	}

	repl(conf)
}
