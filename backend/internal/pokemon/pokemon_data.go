package pokemon

import "github.com/defryheryanto/pokemon-helper/internal/pokemontype"

type PokemonData struct {
	ID         int                 `json:"id"`
	Name       string              `json:"name"`
	BaseStatus *Status             `json:"base_status"`
	Types      []pokemontype.IType `json:"types"`
	Sprites    string              `json:"sprites"`
}
