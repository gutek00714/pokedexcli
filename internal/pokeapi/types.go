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

type Pokemon struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}