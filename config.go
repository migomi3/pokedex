package main

import "github.com/migomi3/pokedex/internal/pokecache"

type Config struct {
	cache           pokecache.Cache
	nextLocationURL *string
	prevLocationURL *string
}
