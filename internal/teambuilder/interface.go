package teambuilder

import "github.com/defry256/pokemon-helper/internal/pokemontype"

type IService interface {
	CalculateTypeCoverage(pokemonNames []string) (coveredTypes, uncoveredTypes []pokemontype.IType, err error)
	CalculateSuggestedType(uncoveredTypes []pokemontype.IType, suggestLength int) []pokemontype.IType
}
