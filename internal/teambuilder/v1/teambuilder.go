package teambuilder

import (
	"context"
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

func (s *Service) CalculateTypeCoverage(ctx context.Context, pokemonNames []string) (typeCovered, typeUncovered []pokemontype.IType, err error) {
	typeCovered = []pokemontype.IType{}
	for _, pokemonName := range pokemonNames {
		pokemonData := s.pokedex.GetPokedex(ctx, pokemonName)
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

func (s *Service) CalculateSuggestedType(ctx context.Context, uncoveredTypes []pokemontype.IType, suggestLength int) []pokemontype.IType {
	suggestedTypeScore := map[pokemontype.IType]int{}

	for key := range getTypesMap() {
		suggestedTypeScore[key] = 0
	}

	for _, uncoveredType := range uncoveredTypes {
		for _, weakAgainst := range uncoveredType.WeakAgainst() {
			suggestedTypeScore[weakAgainst] += 1
		}
	}

	finalSuggestedTypes := []pokemontype.IType{}
	for i := 0; i < suggestLength; i++ {
		var maxType pokemontype.IType
		maxValue := 0
		for key, value := range suggestedTypeScore {
			if value > maxValue {
				maxType = key
				maxValue = value
			}
		}
		if maxType != nil {
			delete(suggestedTypeScore, maxType)
			finalSuggestedTypes = append(finalSuggestedTypes, maxType)
		}
	}

	return finalSuggestedTypes
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
