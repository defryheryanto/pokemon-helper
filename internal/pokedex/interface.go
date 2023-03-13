package pokedex

import (
	"context"

	"github.com/defry256/pokemon-helper/internal/pokemon"
)

type IService interface {
	GetAllPokedex(ctx context.Context, search string) []*pokemon.PokemonData
	GetPokedex(ctx context.Context, pokemonName string) *pokemon.PokemonData
}
