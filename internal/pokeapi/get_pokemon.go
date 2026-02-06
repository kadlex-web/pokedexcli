package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/kadlex-web/pokedexcli/internal/pokecache"
)

func GetPokemon(url string, cache *pokecache.Cache) (PokelistJson, error) {
	// define an empty PokelistJson struct for Marshalling values
	pokemonList := PokelistJson{}

	// perform a conditional checking to see if the passed url value already has been loaded
	val, ok := cache.Get(url)
	// if key value already exists in the cache, unmarshal and return it
	if ok {
		err := json.Unmarshal(val, &pokemonList)
		if err != nil {
			return pokemonList, nil
		}
	} else {
		// otherwise we need to send a request!
		res, err := http.Get(url)
		if err != nil {
			return pokemonList, err
		}
		// reads the response body as bytes
		body, err := io.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			return pokemonList, err
		}
		// add url to cache as a key with the slice of bytes as the body
		cache.Add(url, body)

		_, exists := cache.Get(url)
		if exists {
			fmt.Println("value written to cache")
		}
		// now unmarshal the body and map it to the PokelistJson pokemonList
		err = json.Unmarshal(body, &pokemonList)
		if err != nil {
			return pokemonList, err
		}
		return pokemonList, nil
	}
	return pokemonList, nil
}
