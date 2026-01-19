package pokedex

import (
	"context"

	"github.com/defryheryanto/pokemon-helper/internal/pokemon"
)

type IService interface {
	GetAllPokedex(ctx context.Context, filter GetAllPokedexFilter) ([]pokemon.PokemonData, error)
	GetPokedex(ctx context.Context, pokemonName string) (*pokemon.PokemonData, error)
}
