package pokeapi_tools

import (
	"encoding/json"
	"io"
	"net/http"
)

type locationJson struct {
	// json:count corresponds to the json data representation and allows mapping upon unmarshalling
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocations(url string) (locationJson, error) {
	area_map := locationJson{}
	res, err := http.Get(url)
	if err != nil {
		return area_map, err
	}
	// reads the response body as bytes
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return area_map, err
	}
	// initializes a locationJSON struct for unmarshalling

	err = json.Unmarshal(body, &area_map)
	if err != nil {
		return area_map, err
	}

	return area_map, nil
}
