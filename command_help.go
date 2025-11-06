package main

import (
	"fmt"
	"pokedexcli/internal/pokecache"
)

// callback for help command
func commandHelp(c *config, cache *pokecache.Cache, input []string) error {
	commands := getCommands()

	fmt.Print("Welcome to the Pokedex!\n")
	fmt.Println("Usage:")
	fmt.Println()

	for _, value := range commands {
		fmt.Printf("%s: %s", value.name, value.description)
		fmt.Println()
	}
	return nil
}