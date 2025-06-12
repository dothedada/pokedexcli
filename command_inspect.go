package main

import (
	"fmt"

	"github.com/dothedada/pokemoncli/internals/pokeapi"
)

func commandInspect(conf *config, params ...string) error {
	var pokemonData pokeapi.PokemonData
	var exist bool

	if pokemonData, exist = conf.pokedex[params[0]]; !exist {
		fmt.Printf("You don't have '%s' in your pokedex\n", params[0])
		return nil
	}

	fmt.Println("Name:", pokemonData.Name)
	fmt.Println("Height:", pokemonData.Height)
	fmt.Println("Weight:", pokemonData.Weight)

	if len(pokemonData.Stats) > 0 {
		fmt.Println("Stats:")
		for _, stat := range pokemonData.Stats {
			fmt.Printf("\t- %s: %d\n", stat.Stat.Name, stat.BaseStat)

		}
	}

	if len(pokemonData.Stats) > 0 {
		fmt.Println("Abilities:")
		for _, ability := range pokemonData.Abilities {
			fmt.Println("\t-", ability.Ability.Name)
		}
	}

	if len(pokemonData.Types) > 0 {
		fmt.Println("Types:")
		for _, typeP := range pokemonData.Types {
			fmt.Println("\t-", typeP.Type.Name)
		}
	}

	return nil

}
