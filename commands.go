package main

import (
	"fmt"
	"os"

	"github.com/12awoodward/pokedexcli/internal/pokeapi"
)

type config struct {
	mapPage int
}

type cliCommand struct {
	name string
	description string
	callback func(*config) error
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
		"mapb": {
			name: "map",
			description: "Displays the previous 20 areas",
			callback: commandMapb,
		},
	}
}

func commandMapb(c *config) error {
	if c.mapPage <= 1 {
		fmt.Println("You're on the first page")
		return nil
	}

	c.mapPage -= 1

	areas, err := pokeapi.GetLocationAreas(c.mapPage)
	if err != nil {
		return err
	}
	
	for _, area := range areas {
		fmt.Println(area.Name)
	}
	
	return nil
}

func commandMap(c *config) error {
	c.mapPage += 1

	areas, err := pokeapi.GetLocationAreas(c.mapPage)
	if err != nil {
		return err
	}

	if len(areas) == 0 {
		c.mapPage -= 1
		fmt.Println("You're on the last page")
		return nil
	}
	
	for _, area := range areas {
		fmt.Println(area.Name)
	}
	
	return nil
}

func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, command := range getCommands() {
		fmt.Printf("%v: %v\n", command.name, command.description)
	}
	return nil
}