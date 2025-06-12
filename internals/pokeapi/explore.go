package pokeapi

import (
	"encoding/json"
	"fmt"
)

func (c *Client) ExploreLocation(locationName string) (encountersData, error) {

	url := baseUrl + "/location-area/" + locationName

	var data []byte
	var err error

	data, isCached := c.cache.Get(url)
	if isCached == false {
		data, err = fetchData(url, c.httpClient)
		if err != nil {
			return encountersData{}, fmt.Errorf(
				"Unable to fetch the data for location '%s': %w",
				locationName,
				err,
			)
		}

		c.cache.Add(url, data)
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
