package main

import (
	"fmt"
)

func commandMap(conf *config, param ...string) error {
	mapData, err := conf.client.GetLocations(conf.nextURL)
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

func commandMapBack(conf *config, param ...string) error {
	mapData, err := conf.client.GetLocations(conf.prevURL)
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
