// Structs copied from:
// https://pkg.go.dev/github.com/mtslzr/pokeapi-go@v1.4.0/structs#Pokedex
package pokedex

import (
	"encoding/json"
	"errors"

	_ "embed"
)

var (
	//go:embed pokedex.json
	jsonPokedex []byte

	pokedex Pokedex
)

func init() {
	json.Unmarshal(jsonPokedex, &pokedex)
}

// GetPokedex gets the Pokedex.
func GetPokedex() *Pokedex {
	return &pokedex
}

// GetPokemonEntry returns a single Pokemon entry.
func GetPokemonEntry(id int) (*PokemonEntry, error) {
	if id < 1 || id > len(pokedex.PokemonEntries) {
		return nil, errors.New("invalid id")
	}

	return &pokedex.PokemonEntries[id-1], nil
}

// Pokedex is a single Pokedex.
type Pokedex struct {
	Descriptions []struct {
		Description string `json:"description"`
		Language    struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"descriptions"`
	ID           int    `json:"id"`
	IsMainSeries bool   `json:"is_main_series"`
	Name         string `json:"name"`
	Names        []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEntries []PokemonEntry `json:"pokemon_entries"`
	Region         interface{}    `json:"region"`
	VersionGroups  []interface{}  `json:"version_groups"`
}

// PokemonEntry is a single Pokemon entry.
type PokemonEntry struct {
	EntryNumber    int `json:"entry_number"`
	PokemonSpecies struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"pokemon_species"`
}
