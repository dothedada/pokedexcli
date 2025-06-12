package pokeapi

import (
	"encoding/json"
	"fmt"

	"github.com/dothedada/pokemoncli/internals/pokecache"
)

func (c *Client) GetLocations(
	pageUrl *string,
	cache pokecache.Cache,
) (shallowLocations, error) {

	url := baseUrl + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	var data []byte
	var err error

	data, isCached := cache.Get(url)
	if isCached == false {
		data, err = fetchData(url, c.httpClient)
		if err != nil {
			return shallowLocations{}, err
		}

		cache.Add(url, data)
	}

	var result shallowLocations
	err = json.Unmarshal(data, &result)
	if err != nil {
		return shallowLocations{}, fmt.Errorf(
			"Cannot unmarshal the data: %w",
			err,
		)
	}

	return result, nil
}
