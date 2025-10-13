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
	//takes the cleaned command input and sees if it's in the commands map
	//if there's a match call the call back and return any errors
	elem, ok := getCommands()[command]
	if ok {
		// If the command exists in the commands map, execute it
		err := elem.callback()
		if err != nil {
			fmt.Println(err)
		}
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
	fmt.Print("Welcome to the Pokedex!\n")
	fmt.Println("Usage:")
	fmt.Println()
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
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

// Gets all possible commands. For use in the help function but also allows proper initialization
func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
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
		//fmt.Printf("Your command was %s\n", command)
		ProcessCommand(command)
		fmt.Print("Pokedex > ")
	}
}
