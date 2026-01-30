package main

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/12awoodward/pokedexcli/internal/pokeapi"
	"github.com/12awoodward/pokedexcli/internal/pokecache"
)

type config struct {
	mapNext string
	mapPrev string
	cache pokecache.Cache
	pokedex map[string]pokeapi.Pokemon
}

type cliCommand struct {
	name string
	description string
	callback func(*config, ...string) error
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
		"explore": {
			name: "explore",
			description: "Find all Pokemon at a given area",
			callback: commandExplore,
		},
		"catch": {
			name: "catch",
			description: "Attempt to Catch a given Pokemon",
			callback: commandCatch,
		},
	}
}

func commandCatch(c *config, args ...string) error {
	var pokemon pokeapi.Pokemon
	url := pokeapi.ApiUrl + "pokemon/" + args[0]

	err := getCache(&c.cache, url, &pokemon)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s\n", pokemon.Name)
	
	if rand.Intn(1000) < pokemon.BaseExperience {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	c.pokedex[pokemon.Name] = pokemon
	fmt.Printf("%s was caught!\n", pokemon.Name)

	return nil
}

func commandExplore(c *config, args ...string) error {
	var areaDetails pokeapi.LocationArea
	areaUrl := pokeapi.ApiUrl + "location-area/" + args[0]
	fmt.Printf("Exploring %s...\n", args[0])

	err := getCache(&c.cache, areaUrl, &areaDetails)
	if err != nil {
		return err
	}

	encounters := areaDetails.PokemonEncounters
	if len(encounters) == 0 {
		fmt.Println("No Pokemon Found")
	}

	fmt.Println("Found Pokemon:")
	for _, encounter := range encounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	return nil
}

func commandMapb(c *config, args ...string) error {
	if len(c.mapPrev) == 0 {
		fmt.Println("You're on the first page")
		return nil
	}

	err := getMap(c, c.mapPrev)
	if err != nil {
		return err
	}
	
	return nil
}

func commandMap(c *config, args ...string) error {
	if len(c.mapNext) == 0 {
		c.mapNext = pokeapi.ApiUrl + "location-area"
	}

	err := getMap(c, c.mapNext)
	if err != nil {
		return err
	}
	
	return nil
}

func commandExit(c *config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config, args ...string) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, command := range getCommands() {
		fmt.Printf("%v: %v\n", command.name, command.description)
	}
	return nil
}