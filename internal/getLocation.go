package main

import (
	"fmt"
	"io"
	"net/http"
)

func getLocations() error {
	res, err := http.Get("https://pokeapi.co/api/v2/location-area/")
	if err != nil {
		return err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	fmt.Printf("%s", body)
	return nil
}

func main() {
	getLocations()
}
