package main

import (
	"fmt"
	"strings"
)

func cleanInput(text string) []string {
	// Maybe change order of operations; lowercase first than use fields because fields trims?
	// Step 1 -- trim all leading and trailing whitespace from text string
	trimmedInput := strings.Trim(text)
	// Step 2 -- lowercase text string
	trimmedAndLoweredInput := strings.ToLower(trimmedInput)
	// Step 3 -- split string into a slice of strings, where each item is a word within the original string
	cleanedInput := strings.Fields(trimmedAndLoweredInput)

	return cleanedInput
}

func main() {
	fmt.Println("Hello, World!")
}
