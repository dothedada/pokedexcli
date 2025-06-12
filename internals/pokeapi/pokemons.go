package pokeapi

import (
	"encoding/json"
	"fmt"
)

func (c *Client) GetPokemonData(pokemonName string) (PokemonData, error) {
	url := baseUrl + "/pokemon/" + pokemonName

	var data []byte
	var err error

	data, isCached := c.cache.Get(url)
	if isCached == false {
		data, err = fetchData(url, c.httpClient)
		if err != nil {
			return PokemonData{}, fmt.Errorf(
				"Unable to fetch the data for the Pokemon'%s': %w",
				pokemonName,
				err,
			)
		}

		c.cache.Add(url, data)
	}

	var result PokemonData
	err = json.Unmarshal(data, &result)
	if err != nil {
		return PokemonData{}, fmt.Errorf(
			"Cannot unmarshal the data: %w",
			err,
		)
	}

	return result, nil
}
