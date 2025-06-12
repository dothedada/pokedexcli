package main

import (
	"fmt"
)

func commandExplore(conf *config, params ...string) error {
	if len(params) == 0 {
		return fmt.Errorf("Must provide a location name")
	}

	locationData, err := conf.client.ExploreLocation(params[0])
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
