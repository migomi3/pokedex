package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type LocationAreaRes struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetAreas(url string) (LocationAreaRes, error) {
	res, err := http.Get(url)
	if err != nil {
		return LocationAreaRes{}, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaRes{}, err
	}
	defer res.Body.Close()

	locationArea := LocationAreaRes{}
	err = json.Unmarshal(body, &locationArea)
	if err != nil {
		return LocationAreaRes{}, err
	}

	return locationArea, nil
}
