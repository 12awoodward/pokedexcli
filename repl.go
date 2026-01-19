package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	if len(text) == 0 {
		return []string{}
	}

	split := strings.Fields(text)
	lower := make([]string, len(split))

	for i, word := range split {
		lower[i] = strings.ToLower(word)
	}
	return lower
}

func mainLoop() {
	scanner := bufio.NewScanner(os.Stdin)
	allCommands := getCommands()

	for {
		fmt.Print("Pokedex > ")
		
		scanner.Scan()
		input := scanner.Text()
		
		clean := cleanInput(input)

		if len(clean) > 0 {

			if command, ok := allCommands[clean[0]]; ok {
				command.callback()
			} else {
				fmt.Println("Unknown command")
			}
		}
	}
}