package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/kadlex-web/pokedexcli/internal/pokecache"
)

func GetPokemonData(url string, cache *pokecache.Cache) (PokeDataJson, error) {
	// define an empty PokelistJson struct for Marshalling values
	pokemonData := PokeDataJson{}

	// perform a conditional checking to see if the passed url value already has been loaded
	val, ok := cache.Get(url)
	// if key value already exists in the cache, unmarshal and return it
	if ok {
		err := json.Unmarshal(val, &pokemonData)
		if err != nil {
			return pokemonData, nil
		}
	} else {
		// otherwise we need to send a request!
		res, err := http.Get(url)
		if err != nil {
			return pokemonData, err
		}
		// reads the response body as bytes
		body, err := io.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			return pokemonData, err
		}
		// add url to cache as a key with the slice of bytes as the body
		cache.Add(url, body)

		_, exists := cache.Get(url)
		if exists {
			fmt.Println("value written to cache")
		}
		// now unmarshal the body and map it to the PokelistJson pokemonData
		err = json.Unmarshal(body, &pokemonData)
		if err != nil {
			return pokemonData, err
		}
		return pokemonData, nil
	}
	return pokemonData, nil
}
