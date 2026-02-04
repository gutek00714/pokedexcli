package main

import (
	"github.com/gutek00714/pokedexcli/internal/pokecache"
	"time"
)

func main() {
	myCache := pokecache.NewCache(5 * time.Minute)
	cfg := &config{
		nextLocationsURL: "",
		previousLocationsURL: "",
		pokeCache: myCache,
	}

	startRepl(cfg)
}