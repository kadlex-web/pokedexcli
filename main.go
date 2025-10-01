package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// struct which defines behavior for a cliCommand struct (object)
type cliCommand struct {
	name string
	description string
	callback func() error
}
// map which defines supported commands 
commands := map[string]cliCommand{
	"exit": {
		name: "exit",
		description: "Exit the Pokedex",
		callback: commandExit,
	}
}

// Function which accepts the callback
func ProcessCommand(command string) {
	//takes the cleaned command input and sees if it's in the commands map
	//if there's a match call the call back and return any errors
	elem, ok := commands[command]
	if ok {
		commands[command][callback]
	}
	// if not just print Unknown command
	else {
		fmt.Println("Unknown command")
	}
}

// callback for exit command
func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Pokedex > ")

	for scanner.Scan() {
		input := cleanInput(scanner.Text())
		command := ProcessCommand(input[0])
		fmt.Print("\nPokedex > ")
	}
}
