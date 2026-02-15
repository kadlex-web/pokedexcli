package main

import (
	"fmt"

	"github.com/kadlex-web/pokedexcli/internal/pokecache"
)

func commandInspect(c *config, cache *pokecache.Cache, input []string) error {
	// check if pokemon is in your dex; if not display an error
	name := input[1]
	pokemon, ok := c.pokedex[name]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	// print the following to standard output as the pokemon's info
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	// for every value in the stats sheet -- access struct values and print

	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("\t-%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, poketype := range pokemon.Types {
		fmt.Printf("\t-%v\n", poketype.Type.Name)
	}

	return nil
}
