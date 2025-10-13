package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// struct which defines behavior for a cliCommand struct (object)
type cliCommand struct {
	name        string
	description string
	callback    func() error
}

// Function which accepts the callback
func ProcessCommand(command string) {
	// map which defines supported commands
	commands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:		"help",
			description: "Displays a help message",
			callback:	commandHelp,
		},
	}
	//takes the cleaned command input and sees if it's in the commands map
	//if there's a match call the call back and return any errors
	elem, ok := commands[command]
	if ok {
		// If the command exists in the commands map, execute it 
		elem.callback()
	} else {
		// if not just print Unknown command
		fmt.Println("Unknown command")
	}
}

// callback for exit command
func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
// callback for help command
func commandHelp() error {
	println("asking for help I see..")
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


func startRepl() {
	// need to gracefully handle if the user submits a command with no input
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Pokedex > ")
	for scanner.Scan() {
		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			fmt.Println("Command of 0 length submitted")
		}
		command := input[0]
		fmt.Printf("Your command was %s\n", command)
		ProcessCommand(command)
		fmt.Print("Pokedex > ")
	}
}