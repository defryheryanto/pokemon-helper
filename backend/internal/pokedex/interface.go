package pokedex

//go:generate go run github.com/golang/mock/mockgen -source=interface.go -package=mock -destination=mock/mock.go

import (
	"context"

	"github.com/defryheryanto/pokemon-helper/internal/pokemon"
)

type IService interface {
	GetAllPokedex(ctx context.Context, filter GetAllPokedexFilter) ([]pokemon.PokemonData, error)
	GetPokedex(ctx context.Context, pokemonName string) (*pokemon.PokemonData, error)
}
