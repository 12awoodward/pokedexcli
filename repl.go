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
		input := cleanInput(scanner.Text())

		if len(input) > 0 {

			if command, ok := allCommands[input[0]]; ok {

				err := command.callback()
				if err != nil {
					fmt.Printf("Error %v\n", err)
				}

			} else {
				fmt.Println("Unknown command")
			}
		}
	}
}