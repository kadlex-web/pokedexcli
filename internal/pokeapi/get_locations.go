package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/kadlex-web/pokedexcli/internal/pokecache"
)

func GetLocations(url string, cache *pokecache.Cache) (LocationJson, error) {
	// define an empty LocationJson struct for Marshalling values
	area_map := LocationJson{}

	// perform a conditional checking to see if the passed url value already has been loaded
	val, ok := cache.Get(url)
	// if key value already exists in the cache, unmarshal and return it
	if ok {
		err := json.Unmarshal(val, &area_map)
		if err != nil {
			return area_map, nil
		}
	} else {
		// otherwise we need to send a request!
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

		_, exists := cache.Get(url)
		if exists {
			fmt.Println("value written to cache")
		}
		// now unmarshal the body and map it to the LocationJson area_map
		err = json.Unmarshal(body, &area_map)
		if err != nil {
			return area_map, err
		}
		return area_map, nil
	}
	return area_map, nil
}
