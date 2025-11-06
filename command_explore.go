package main

import (
	"fmt"
	"pokedexcli/internal/pokecache"
	"net/http"
)

func commandExplore(c *config, cache *pokecache.Cache, input []string) error {
	res, err := http.Get("http://google.com")
	if err != nil {
		return err
	}
	fmt.Println(res)
	fmt.Println("not yet implemented")
	return nil
}