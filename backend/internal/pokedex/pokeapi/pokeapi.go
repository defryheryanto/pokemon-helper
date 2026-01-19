package pokeapi

import (
	"context"

	pokego "github.com/JoshGuarino/PokeGo/pkg"
	"github.com/defryheryanto/pokemon-helper/internal/pokedex"
	"github.com/defryheryanto/pokemon-helper/internal/pokemon"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type PokeAPI struct {
	client     pokego.PokeGo
	titleCaser cases.Caser
}

func NewPokeAPI() *PokeAPI {
	client := pokego.NewClient()
	titleCaser := cases.Title(language.English)
	return &PokeAPI{
		client:     client,
		titleCaser: titleCaser,
	}
}

func (p *PokeAPI) GetAllPokedex(ctx context.Context, filter pokedex.GetAllPokedexFilter) ([]pokemon.PokemonData, error) {
	pokemons, err := p.client.Pokemon.GetPokemonList(filter.PageSize, (filter.Page-1)*filter.PageSize)
	if err != nil {
		return nil, err
	}

	result := make([]pokemon.PokemonData, 0, len(pokemons.Results))
	for _, poke := range pokemons.Results {
		result = append(result, pokemon.PokemonData{
			Name: p.titleCaser.String(poke.Name),
		})
	}

	return result, nil
}

func (p *PokeAPI) GetPokedex(ctx context.Context, pokemonName string) (*pokemon.PokemonData, error) {
	poke, err := p.client.Pokemon.GetPokemon(pokemonName)
	if err != nil {
		return nil, err
	}

	return convertToPokemonData(poke), nil
}
