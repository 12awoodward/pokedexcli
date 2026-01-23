package main

import (
	"fmt"
	"os"

	"github.com/12awoodward/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name string
	description string
	callback func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"map": {
			name: "map",
			description: "Displays the next 20 areas",
			callback: commandMap,
		},
	}
}

func commandMap() error {
	areas, _, err := pokeapi.GetLocationAreas(100)
	if err != nil {
		return err
	}
	
	for _, area := range areas {
		fmt.Println(area.Name)
	}
	
	return nil
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, command := range getCommands() {
		fmt.Printf("%v: %v\n", command.name, command.description)
	}
	return nil
}