package pokeapi

import (
	"fmt"
	"io"
	"net/http"
)

func fetchData(url string, client http.Client) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []byte{}, fmt.Errorf(
			"Cannot create the http request: %w",
			err,
		)
	}

	res, err := client.Do(req)
	if err != nil {
		return []byte{}, fmt.Errorf(
			"Cannot create the response element: %w",
			err,
		)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, fmt.Errorf(
			"Cannot create the data element: %w",
			err,
		)
	}
	return data, nil
}
