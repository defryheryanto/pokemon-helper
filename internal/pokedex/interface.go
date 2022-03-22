package pokedex

import "github.com/defry256/pokemon-helper/internal/pokemon"

type IService interface {
	GetAllPokedex(search string) []*pokemon.PokemonData
	GetPokedex(pokemonName string) *pokemon.PokemonData
}
