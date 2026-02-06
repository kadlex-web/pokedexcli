package pokeapi

type LocationJson struct {
	// json:count corresponds to the json data representation and allows mapping upon unmarshalling
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
