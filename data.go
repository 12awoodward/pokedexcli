package main

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/12awoodward/pokedexcli/internal/pokeapi"
	"github.com/12awoodward/pokedexcli/internal/pokecache"
)

func getMap(c *config, url string) error {
	var areas pokeapi.LocationAreas

	err := getCache(&c.cache, url, &areas)
	if err != nil {
		return err
	}

	c.mapNext = areas.Next
	c.mapPrev = areas.Previous
	
	for _, area := range areas.Results {
		fmt.Println(area.Name)
	}

	return nil
}

func getCache[T any](c *pokecache.Cache, url string, dest *T) error {
	data, ok := c.Get(url)
	if ok {
		decoder := gob.NewDecoder(bytes.NewReader(data))
		err := decoder.Decode(dest)
		if err != nil {
			return err
		}
		return nil
	}

	err := pokeapi.GetApiData(url, dest)
	if err != nil {
		return err
	}

	buffer := bytes.Buffer{}
	encoder := gob.NewEncoder(&buffer)
	err = encoder.Encode(*dest)
	if err != nil {
		return err
	}

	c.Add(url, buffer.Bytes())
	return nil
}