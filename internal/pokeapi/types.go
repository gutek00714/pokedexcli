package pokeapi

type Issue struct {
	Next string `json:"next"`
	Previous string `json:"previous"`
	Results []Result `json:"results"`
}

type Result struct {
	Name string `json:"name"`
}

type LocationArea struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"` 
	} `json:"pokemon_encounters"`	
}