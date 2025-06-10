package main

import (
	"time"

	"github.com/dothedada/pokemoncli/internals/pokeapi"
)

func main() {
	client := pokeapi.NewClient(3 * time.Second)
	conf := &config{client: client}

	repl(conf)
}
