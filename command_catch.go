package main

import (
	"fmt"
	"math/rand"

	"github.com/kadlex-web/pokedexcli/internal/pokeapi"
	"github.com/kadlex-web/pokedexcli/internal/pokecache"
)

func catch(exp int) bool {
	fmt.Println("Pokemon's base exp is:", exp)
	// x is a value between 0 and base exp
	x := rand.Intn(exp)
	fmt.Println("value is", x)
	// if random value is greater than 50% of the base exp you catch it
	if x >= exp/2 {
		return true
	}
	return false
}

func commandCatch(c *config, cache *pokecache.Cache, input []string) error {
	baseUrl := "https://pokeapi.co/api/v2/pokemon/"
	mon := input[1]
	url := baseUrl + mon
	pokeData, err := pokeapi.GetPokemonData(url, cache)
	if err != nil {
		fmt.Println("Requested Pokemon does not exist")
		return nil
	} else {
		if pokeData.BaseExperience == 0 {
			fmt.Println("Requested Pokemon does not exist!")
			return nil
		}
		fmt.Printf("Throwing a Pokeball at %v\n", mon)
		if catch(pokeData.BaseExperience) {
			fmt.Printf("%v was caught\n", mon)
			//TODO: add pokemon to pokedex
		} else {
			fmt.Printf("%v escaped!\n", mon)
		}
	}

	return nil
}
