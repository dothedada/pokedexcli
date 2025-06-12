package main

import "fmt"

func commandPokedex(conf *config, args ...string) error {
	if len(conf.pokedex) == 0 {
		fmt.Println("You havent caught any Pokemon :(")
	}

	fmt.Println("Your Pokedex has:")
	for key := range conf.pokedex {
		fmt.Println(" -", key)
	}

	return nil
}
