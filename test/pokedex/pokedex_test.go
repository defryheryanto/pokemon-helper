package pokedex_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/defry256/pokemon-helper/internal/pokedex/v1"
	"github.com/defry256/pokemon-helper/internal/pokemon"
	"github.com/defry256/pokemon-helper/internal/pokemontype"
)

func TestGetAllPokemonData(test *testing.T) {
	service := pokedex.NewService()
	pokemons := service.GetAllPokedex(context.Background(), "")
	for _, poke := range pokemons {
		if poke.Name == "" {
			test.Fatalf("Pokemon name is empty")
		}
	}
}

func TestSearchPokedex(test *testing.T) {
	service := pokedex.NewService()
	pokemons := service.GetAllPokedex(context.Background(), "bulba")
	if pokemons[0].Name != "Bulbasaur" {
		test.Fatalf(fmt.Sprintf("Pokemon Name Expected: Bulbasaur, Got: %s", pokemons[0].Name))
	}
}

func TestGetPokemonData(test *testing.T) {
	pokemonName := "geodude"
	expectedData := &pokemon.PokemonData{
		Name: "Geodude",
		BaseStatus: &pokemon.Status{
			HP:             40,
			Attack:         80,
			Defense:        100,
			SpecialAttack:  30,
			SpecialDefense: 30,
			Speed:          20,
			Total:          300,
		},
		Types: []pokemontype.IType{
			pokemontype.RockType,
			pokemontype.GroundType,
		},
	}

	service := pokedex.NewService()
	pokemonData := service.GetPokedex(context.TODO(), pokemonName)
	if !pokemonDataEqual(expectedData, pokemonData) {
		test.Fatalf(fmt.Sprintf("Actual pokemon data and expected data not equal: %v != %v", pokemonData, expectedData))
	}
}

func pokemonDataEqual(data1, data2 *pokemon.PokemonData) bool {
	if data1.Name != data2.Name ||
		data1.BaseStatus.Attack != data2.BaseStatus.Attack ||
		data1.BaseStatus.Defense != data2.BaseStatus.Defense ||
		data1.BaseStatus.HP != data2.BaseStatus.HP ||
		data1.BaseStatus.SpecialAttack != data2.BaseStatus.SpecialAttack ||
		data1.BaseStatus.SpecialDefense != data2.BaseStatus.SpecialDefense ||
		data1.BaseStatus.Speed != data2.BaseStatus.Speed ||
		data1.BaseStatus.Total != data2.BaseStatus.Total {
		return false
	}
	if len(data1.Types) != len(data2.Types) {
		return false
	}
	for i, t := range data1.Types {
		if t != data2.Types[i] {
			return false
		}
	}

	return true
}
