package pokeapi

import (
	"errors"
	"net/url"
	"strconv"
)

type locationAreas struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []locationAreasResult `json:"results"`
}

type locationAreasResult struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func GetLocationAreas(pageNo int) ([]locationAreasResult, error) {
	if pageNo < 0 {
		return nil, errors.New("page cannot be negative")
	}
	offset := strconv.Itoa(pageNo * 20)

	query := url.Values{
		"limit": {"20"},
		"offset": {offset},
	}

	requestUrl, err := addQueryParams(apiUrl + "location-area", query)
	if err != nil {
		return nil, err
	}

	var areas locationAreas
	err = getApiData(requestUrl, &areas)
	if err != nil {
		return nil, err
	}

	return areas.Results, nil
}