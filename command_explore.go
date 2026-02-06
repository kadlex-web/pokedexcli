package main

import (
	"fmt"

	"github.com/kadlex-web/pokedexcli/internal/pokeapi"
	"github.com/kadlex-web/pokedexcli/internal/pokecache"
)

func commandExplore(c *config, cache *pokecache.Cache, input []string) error {
	location := input[1]
	url := "https://pokeapi.co/api/v2/location-area/" + location
	pokeList, _ := pokeapi.GetPokemon(url, cache)
	if len(pokeList.PokemonEncounters) == 0 {
		fmt.Println("Invalid Location or No Pokemon Found Here!")
		fmt.Println("Try a different location!")
	}
	for _, pokemon := range pokeList.PokemonEncounters {
		fmt.Printf("%v\n", pokemon.Pokemon.Name)
	}
	return nil
}
