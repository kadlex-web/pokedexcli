package main

import (
	"fmt"

	"github.com/kadlex-web/pokedexcli/internal/pokecache"
)

func commandPokedex(c *config, cache *pokecache.Cache, input []string) error {
	if len(c.pokedex) <= 0 {
		fmt.Println("Your pokedex is empty!")
		return nil
	}
	for _, pokemon := range c.pokedex {
		fmt.Println("  -", pokemon.Name)
	}
	return nil
}
