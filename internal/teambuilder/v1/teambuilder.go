package teambuilder

import (
	"fmt"

	"github.com/defry256/pokemon-helper/internal/errors"
	"github.com/defry256/pokemon-helper/internal/pokedex"
	"github.com/defry256/pokemon-helper/internal/pokemontype"
)

type Service struct {
	pokedex pokedex.IService
}

func NewService(pokedexService pokedex.IService) *Service {
	return &Service{pokedexService}
}

func (s *Service) CalculateTypeCoverage(pokemonNames []string) (typeCovered, typeUncovered []pokemontype.IType, err error) {
	typeCovered = []pokemontype.IType{}
	for _, pokemonName := range pokemonNames {
		pokemonData := s.pokedex.GetPokedex(pokemonName)
		if pokemonData == nil {
			return nil, nil, errors.NewNotFoundError(fmt.Sprintf("Pokemon %s not found", pokemonName))
		}
		for _, pokemonType := range pokemonData.Types {
			typeCovered = append(typeCovered, pokemonType.StrongAgainst()...)
		}
	}

	typeUncovered = getUncoveredType(typeCovered)

	return typeCovered, typeUncovered, err
}

func getUncoveredType(typeCovered []pokemontype.IType) []pokemontype.IType {
	uncoveredMap := getTypesMap()

	for _, t := range typeCovered {
		_, ok := uncoveredMap[t]
		if ok {
			delete(uncoveredMap, t)
		}
	}

	uncoveredTypes := []pokemontype.IType{}
	for _, v := range uncoveredMap {
		uncoveredTypes = append(uncoveredTypes, v)
	}

	return uncoveredTypes
}

func getTypesMap() map[pokemontype.IType]pokemontype.IType {
	types := map[pokemontype.IType]pokemontype.IType{
		pokemontype.NormalType:   pokemontype.NormalType,
		pokemontype.FireType:     pokemontype.FireType,
		pokemontype.WaterType:    pokemontype.WaterType,
		pokemontype.ElectricType: pokemontype.ElectricType,
		pokemontype.GrassType:    pokemontype.GrassType,
		pokemontype.IceType:      pokemontype.IceType,
		pokemontype.FightingType: pokemontype.FightingType,
		pokemontype.PoisonType:   pokemontype.PoisonType,
		pokemontype.GroundType:   pokemontype.GroundType,
		pokemontype.FlyingType:   pokemontype.FlyingType,
		pokemontype.PsychicType:  pokemontype.PsychicType,
		pokemontype.BugType:      pokemontype.BugType,
		pokemontype.RockType:     pokemontype.RockType,
		pokemontype.GhostType:    pokemontype.GhostType,
		pokemontype.DragonType:   pokemontype.DragonType,
		pokemontype.DarkType:     pokemontype.DarkType,
		pokemontype.SteelType:    pokemontype.SteelType,
		pokemontype.FairyType:    pokemontype.FairyType,
	}
	return types
}