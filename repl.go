package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"pokedexcli/internal/pokeapi_tools"
	"pokedexcli/internal/pokecache"
	"strings"
	"time"
)

// struct which defines behavior for a cliCommand struct (object)
type cliCommand struct {
	name        string
	description string
	callback    func(*config, *pokecache.Cache) error
}

// config for storing previous and next map urls
type config struct {
	next_url string
	previous_url string
}

// map which defines supported commands
func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the name of 20 locations on the overworld",
			callback:    commandMap,
		},
		"mapb": {
			name:		"mapb",
			description: "displays the previous 20 locations on the overworld",
			callback: commandMapBack,
		},
	}
}

// Function which accepts the callback
func ProcessCommand(command string, c *config, cache *pokecache.Cache) {
	//takes the cleaned command input and sees if it's in the commands map
	//if there's a match call the call back and return any errors
	commands := getCommands()
	elem, ok := commands[command]
	if ok {
		// If the command exists in the commands map, execute it
		elem.callback(c, cache)
	} else {
		// if not just print Unknown command
		fmt.Println("Unknown command")
	}
}

// callback for exit command
func commandExit(c *config, cache *pokecache.Cache) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMap(c *config, cache *pokecache.Cache) error {
	// initial state -- nothing in the cache, first map command fired
	if (c.previous_url == "" && c.next_url == "") {
		current_map, _ := pokeapi_tools.GetLocations("https://pokeapi.co/api/v2/location-area/", cache)

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

func commandMapBack(c *config, cache *pokecache.Cache) error {
	// if user is on the first page, give them feedback
	if (c.previous_url == "") {
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

// callback for help command
func commandHelp(c *config, cache *pokecache.Cache) error {
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

// Cleans the input and returns a []string of the cleaned inputs
func cleanInput(text string) []string {
	// Maybe change order of operations; lowercase first than use fields because fields trims?
	// Step 1 -- trim all leading and trailing whitespace from text string
	trimmedInput := strings.Trim(text, " ")
	// Step 2 -- lowercase text string
	trimmedAndLoweredInput := strings.ToLower(trimmedInput)
	// Step 3 -- split string into a slice of strings, where each item is a word within the original string
	cleanedInput := strings.Fields(trimmedAndLoweredInput)
	return cleanedInput
}

func startRepl() {
	// need to gracefully handle if the user submits a command with no input
	scanner := bufio.NewScanner(os.Stdin)
	// initialize an empty configuration file
	c := config {
		next_url: "",
		previous_url: "",
	}
	// initialize a new cache
	cache := pokecache.NewCache(5* time.Second)

	fmt.Print("Pokedex > ")
	for scanner.Scan() {
		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			fmt.Println("Command of 0 length submitted")
		}
		command := input[0]
		//fmt.Printf("Your command was %s\n", command)
		ProcessCommand(command, &c, &cache)
		fmt.Print("Pokedex > ")
	}
}
