package main

import (
	"fmt"
)

func commandExplore(conf *config, param ...string) error {

	locationData, err := conf.client.ExploreLocation(param[0], conf.cache)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", locationData.Name)
	fmt.Println("Found Pokemon:")
	for _, encounter := range locationData.PokemonEncounters {
		fmt.Println(" -", encounter.Pokemon.Name)
	}

	return nil
}
