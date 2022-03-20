package pokemon

import "github.com/defry256/pokemon-helper/internal/pokemontype"

type PokemonData struct {
	Name       string              `json:"name"`
	BaseStatus *Status             `json:"base_status"`
	Types      []pokemontype.IType `json:"types"`
}
