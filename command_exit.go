package main

import (
	"fmt"
	"os"
	"pokedexcli/internal/pokecache"
)

// callback for exit command
func commandExit(c *config, cache *pokecache.Cache, input []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
