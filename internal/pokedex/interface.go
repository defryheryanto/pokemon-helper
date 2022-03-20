package pokedex

import "github.com/defry256/pokemon-helper/internal/pokemon"

type IService interface {
	GetPokedex(pokemonName string) *pokemon.PokemonData
}
