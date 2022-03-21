package teambuilder

import (
	"github.com/defry256/pokemon-helper/internal/pokemon"
	"github.com/defry256/pokemon-helper/internal/pokemontype"
)

type simulateTeamResponse struct {
	Pokemons       []*pokemon.PokemonData `json:"pokemons"`
	CoveredTypes   []pokemontype.IType    `json:"effective_types"`
	UncoveredTypes []pokemontype.IType    `json:"uncovered_types"`
}
