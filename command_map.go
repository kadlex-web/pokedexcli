package main

import (
	"pokedexcli/internal/pokecache"
	"pokedexcli/internal/pokeapi_tools"
	"fmt"
)

func commandMap(c *config, cache *pokecache.Cache, input []string) error {
	// initial state -- nothing in the cache, first map command fired
	if c.previous_url == "" && c.next_url == "" {
		current_map, _ := pokeapi_tools.GetLocations("https://pokeapi.co/api/v2/location-area/?offset=0&limit=20", cache)

		// prints all the results within the json object
		for _, location := range current_map.Results {
			fmt.Printf("%v\n", location.Name)
		}
		c.next_url = current_map.Next
		c.previous_url = "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"
	} else {
		current_map, _ := pokeapi_tools.GetLocations(c.next_url, cache)
		for _, location := range current_map.Results {
			fmt.Printf("%v\n", location.Name)
		}
		c.next_url = current_map.Next
		c.previous_url = current_map.Previous
	}
	return nil
}

func commandMapBack(c *config, cache *pokecache.Cache, input []string) error {
	// if user is on the first page, give them feedback
	if c.previous_url == "" {
		fmt.Println("youre already on the first page")
		return nil
	}
	current_map, _ := pokeapi_tools.GetLocations(c.previous_url, cache)
	for _, location := range current_map.Results {
		fmt.Printf("%v\n", location.Name)
	}
	c.next_url = current_map.Next
	c.previous_url = current_map.Previous
	return nil
}