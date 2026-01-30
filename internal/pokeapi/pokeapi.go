package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const ApiUrl = "https://pokeapi.co/api/v2/"

func GetApiData[T any](endpoint string, result *T) error {
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return err
	}

	client := http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return fmt.Errorf("Status: %v, %v", res.StatusCode, http.StatusText(res.StatusCode))
	}

	decoder := json.NewDecoder(res.Body)
	decoder.Decode(result)
	return nil
}

func AddQueryParams(req string, vals url.Values) (string, error) {
	requestUrl, err := url.Parse(req)
	if err != nil {
		return "", err
	}

	requestUrl.RawQuery = vals.Encode()
	return requestUrl.String(), nil
}