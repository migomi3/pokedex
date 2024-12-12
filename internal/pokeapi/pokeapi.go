package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/migomi3/pokedex/internal/pokecache"
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

func GetAreas(url string, cache *pokecache.Cache) (LocationAreaRes, error) {
	res, err := http.Get(url)
	if err != nil {
		return LocationAreaRes{}, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaRes{}, err
	}
	defer res.Body.Close()

	cache.Add(url, body)

	return UnmarshalAreas(body)
}

func UnmarshalAreas(body []byte) (LocationAreaRes, error) {
	locationArea := LocationAreaRes{}
	err := json.Unmarshal(body, &locationArea)
	if err != nil {
		return LocationAreaRes{}, err
	}

	return locationArea, nil
}
