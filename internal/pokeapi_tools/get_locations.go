package pokeapi_tools

import (
	"encoding/json"
	"io"
	"fmt"
	"net/http"
	"pokedexcli/internal/pokecache"
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

func GetLocations(url string, cache *pokecache.Cache) (locationJson, error) {
	// define an empty locationJson struct for Marshalling values
	area_map := locationJson{}

	// perform a conditional checking to see if the passed url value already has been loaded
	val, ok := cache.Get(url)
	// if key value already exists in the cache, unmarshal and return it
	if ok {
		fmt.Println("using the cache:", val) //logging note
		err := json.Unmarshal(val, &area_map)
		if err != nil {
			return area_map, nil
		}
	}
	// otherwise we need to send a request!
	fmt.Println("not using the cache -- sending an API request") //logging note
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
	// add url to cache as a key with the slice of bytes as the body
	cache.Add(url, body)

	// now unmarshal the body and map it to the locationJSON area_map
	err = json.Unmarshal(body, &area_map)
	if err != nil {
		return area_map, err
	}

	return area_map, nil
}
