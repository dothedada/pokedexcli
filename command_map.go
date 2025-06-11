package main

import (
	"fmt"
)

func commandMap(conf *config) error {
	mapData, err := conf.client.GetShallowLocations(conf.nextURL, conf.cache)
	if err != nil {
		return fmt.Errorf("cannot get the locations: %w", err)
	}

	for _, locationName := range mapData.Results {
		fmt.Println(locationName.Name)
	}

	conf.nextURL = &mapData.Next
	conf.prevURL = &mapData.Previous

	return nil
}

func commandMapBack(conf *config) error {
	mapData, err := conf.client.GetShallowLocations(conf.prevURL, conf.cache)
	if err != nil {
		return fmt.Errorf("cannot get the locations: %w", err)
	}

	for _, locationName := range mapData.Results {
		fmt.Println(locationName.Name)
	}

	conf.nextURL = &mapData.Next
	conf.prevURL = &mapData.Previous

	return nil
}
