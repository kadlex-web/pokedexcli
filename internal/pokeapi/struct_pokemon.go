package pokeapi

// define pokemon struct which unmarshals the JSON response into a pokemon object to be stored in a map
type Pokemon struct {
	Name   string
	Height int
	Weight int
}
