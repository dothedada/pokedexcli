package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/dothedada/pokemoncli/internals/pokecache"
)

func (c *Client) GetShallowLocations(
	pageUrl *string,
	cache pokecache.Cache,
) (shallowLocations, error) {

	url := baseUrl + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	data, isCached := cache.Get(url)
	if isCached == false {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return shallowLocations{}, fmt.Errorf(
				"Cannot create the http request: %w",
				err,
			)
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return shallowLocations{}, fmt.Errorf(
				"Cannot create the response element: %w",
				err,
			)
		}
		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return shallowLocations{}, fmt.Errorf(
				"Cannot create the data element: %w",
				err,
			)
		}

		fmt.Println("NO EN CACHE")

		cache.Add(url, data)
	}

	var result shallowLocations
	err := json.Unmarshal(data, &result)
	if err != nil {
		return shallowLocations{}, fmt.Errorf(
			"Cannot unmarshal the data: %w",
			err,
		)
	}

	return result, nil
}
