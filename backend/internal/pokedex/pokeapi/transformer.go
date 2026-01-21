package pokeapi

import (
	pokegomodels "github.com/JoshGuarino/PokeGo/pkg/models"
	"github.com/defryheryanto/pokemon-helper/internal/pokemon"
	"github.com/defryheryanto/pokemon-helper/internal/pokemontype"
	"golang.org/x/text/cases"
)

func convertToPokemonData(poke *pokegomodels.Pokemon, nameCaser cases.Caser) *pokemon.PokemonData {
	if poke == nil {
		return nil
	}

	stats := &pokemon.Status{}
	for _, stat := range poke.Stats {
		switch stat.Stat.Name {
		case "hp":
			stats.HP = stat.BaseStat
		case "attack":
			stats.Attack = stat.BaseStat
		case "defense":
			stats.Defense = stat.BaseStat
		case "special-attack":
			stats.SpecialAttack = stat.BaseStat
		case "special-defense":
			stats.SpecialDefense = stat.BaseStat
		case "speed":
			stats.Speed = stat.BaseStat
		}
		stats.Total += stat.BaseStat
	}

	types := []pokemontype.IType{}
	for _, t := range poke.Types {
		types = append(types, pokemontype.Type(t.Type.Name))
	}

	return &pokemon.PokemonData{
		ID:         poke.ID,
		Name:       nameCaser.String(poke.Name),
		BaseStatus: stats,
		Types:      types,
		Sprites:    poke.Sprites.FrontDefault,
	}
}
