package pokeapi

import (
	"encoding/json"
	"fmt"
)

func (c *Client) GetLocations(pageUrl *string) (shallowLocations, error) {

	url := baseUrl + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	var data []byte
	var err error

	data, isCached := c.cache.Get(url)
	if isCached == false {
		data, err = fetchData(url, c.httpClient)
		if err != nil {
			return shallowLocations{}, err
		}

		c.cache.Add(url, data)
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
