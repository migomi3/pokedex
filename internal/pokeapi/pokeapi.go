package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/migomi3/pokedex/internal/pokecache"
)

func GetLocationAreasRes(url *string, cache *pokecache.Cache) (LocationAreaRes, error) {
	res, err := http.Get(*url)
	if err != nil {
		return LocationAreaRes{}, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaRes{}, err
	}
	defer res.Body.Close()

	cache.Add(*url, body)

	return UnmarshalLocationAreasRes(body)
}

func UnmarshalLocationAreasRes(body []byte) (LocationAreaRes, error) {
	locationArea := LocationAreaRes{}
	err := json.Unmarshal(body, &locationArea)
	if err != nil {
		return LocationAreaRes{}, err
	}

	return locationArea, nil
}

func GetLocationArea(url *string, cache *pokecache.Cache) (LocationArea, error) {
	res, err := http.Get(*url)
	if err != nil {
		return LocationArea{}, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, err
	}
	defer res.Body.Close()

	cache.Add(*url, body)

	return UnmarshalLocationArea(body)
}

func UnmarshalLocationArea(body []byte) (LocationArea, error) {
	locationArea := LocationArea{}
	err := json.Unmarshal(body, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}

	return locationArea, nil
}
