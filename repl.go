package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/kadlex-web/pokedexcli/internal/pokecache"
)

// struct which defines behavior for a cliCommand struct (object)
type cliCommand struct {
	name        string
	description string
	callback    func(*config, *pokecache.Cache, []string) error
}

// config for storing previous and next map urls
type config struct {
	next_url     string
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
			name:        "mapb",
			description: "displays the previous 20 locations on the overworld",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore",
			description: "lists the pokemon located in the area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "attempts to catch a pokemon",
			callback:    commandCatch,
		},
	}
}

// Function which accepts the callback
func ProcessCommand(input []string, c *config, cache *pokecache.Cache) {
	//takes the cleaned command input and sees if it's in the commands map
	//if there's a match call the call back and return any errors
	commands := getCommands()
	// TODO: Rewrite commands to accept multiple inputs. for exploring locations or catching pokemon
	elem, ok := commands[input[0]]
	if ok {
		elem.callback(c, cache, input)
	} else {
		fmt.Println("unknown command")
	}
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
	c := config{
		next_url:     "",
		previous_url: "",
	}
	// initialize a new cache
	cache := pokecache.NewCache(15 * time.Second)

	fmt.Print("Pokedex > ")
	for scanner.Scan() {
		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			fmt.Println("Command of 0 length submitted")
		}
		ProcessCommand(input, &c, &cache)
		fmt.Print("Pokedex > ")
	}
}
