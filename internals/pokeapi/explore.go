package pokeapi

import (
	"encoding/json"
	"fmt"

	"github.com/dothedada/pokemoncli/internals/pokecache"
)

func (c *Client) ExploreLocation(
	locationName string,
	cache pokecache.Cache,
) (encountersData, error) {

	if locationName == "" {
		return encountersData{}, fmt.Errorf("Must provide the locatio name")
	}
	url := baseUrl + "/location-area/" + locationName

	var data []byte
	var err error

	data, isCached := cache.Get(url)
	if isCached == false {
		data, err = fetchData(url, c.httpClient)
		if err != nil {
			return encountersData{}, fmt.Errorf(
				"Unable to fetch the data for location '%s': %w",
				locationName,
				err,
			)
		}

		cache.Add(url, data)
	}

	var result encountersData
	err = json.Unmarshal(data, &result)
	if err != nil {
		return encountersData{}, fmt.Errorf(
			"Cannot unmarshal the data: %w",
			err,
		)
	}

	return result, nil
}
