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

	for i, word := range split {
		split[i] = strings.ToLower(word)
	}
	return split
}

func mainLoop() {
	scanner := bufio.NewScanner(os.Stdin)
	allCommands := getCommands()

	c := config{}

	for {
		fmt.Print("Pokedex > ")
		
		scanner.Scan()
		input := cleanInput(scanner.Text())

		if len(input) > 0 {

			if command, ok := allCommands[input[0]]; ok {

				err := command.callback(&c)
				if err != nil {
					fmt.Printf("Error %v\n", err)
				}

			} else {
				fmt.Println("Unknown command")
			}
		}
	}
}