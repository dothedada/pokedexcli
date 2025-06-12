package main

import (
	"fmt"
	"math"
	"math/rand"
)

const (
	minXP = 36  // minimum base xperience regitered in pokeapi
	maxXP = 608 // maximum base xperience regitered in pokeapi

	// probability parameters
	amplitude = 0.85 // maximum % of diference between minXP and maxXP
	base      = 0.05 // minimum probability of catching a Pokemon
	fadeRate  = 4.5  // sptepness of the curve (4.5 or higher)
)

func commandCatch(conf *config, params ...string) error {
	if len(params) == 0 {
		return fmt.Errorf("Must provide a Pokemon name")
	}

	pokemonData, err := conf.client.GetPokemonData(params[0])
	if err != nil {
		return err
	}

	if _, exist := conf.pokedex[pokemonData.Name]; exist {
		fmt.Printf("%s is already in your pokedex!\n", pokemonData.Name)
		return nil
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonData.Name)
	if catchTryValue() < catchProbability(pokemonData.BaseExperience) {
		fmt.Printf("%s was caught!\n", pokemonData.Name)
		conf.pokedex[pokemonData.Name] = pokemonData
	} else {
		fmt.Printf("%s escaped!\n", pokemonData.Name)
	}

	return nil

}

func catchProbability(pokemonXP int) float64 {
	normalizedValues := float64(pokemonXP-minXP) / float64(maxXP-minXP)
	return amplitude*math.Exp(-fadeRate*normalizedValues) + base
}

func catchTryValue() float64 {
	return rand.Float64()
}
