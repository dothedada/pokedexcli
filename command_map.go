package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type fetchedLocations struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

var locationUrl = struct {
	previous string
	current  string
	next     string
}{
	previous: "https://pokeapi.co/api/v2/location-area/",
	next:     "https://pokeapi.co/api/v2/location-area/",
}

func commandMap() error {
	res, err := http.Get(locationUrl.next)
	if err != nil {
		fmt.Println("error fetching the data")
	}
	defer res.Body.Close()

	locations := fetchedLocations{}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("cannot load data")
	}
	json.Unmarshal(body, &locations)

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	if locations.Previous != nil {
		locationUrl.previous = locations.Previous.(string)
	}
	locationUrl.current = locationUrl.next
	locationUrl.next = locations.Next

	return nil
}

func commandMapBack() error {
	res, err := http.Get(locationUrl.previous)
	if err != nil {
		fmt.Println("error fetching the data")
	}
	defer res.Body.Close()

	locations := fetchedLocations{}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("cannot load data")
	}
	json.Unmarshal(body, &locations)

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	if locations.Previous != nil {
		locationUrl.previous = locations.Previous.(string)
	}
	locationUrl.current = locationUrl.previous
	locationUrl.next = locations.Next

	return nil
}
