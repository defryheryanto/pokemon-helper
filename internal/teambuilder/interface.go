package teambuilder

import (
	"context"

	"github.com/defry256/pokemon-helper/internal/pokemontype"
)

type IService interface {
	CalculateTypeCoverage(ctx context.Context, pokemonNames []string) (coveredTypes, uncoveredTypes []pokemontype.IType, err error)
	CalculateSuggestedType(ctx context.Context, uncoveredTypes []pokemontype.IType, suggestLength int) []pokemontype.IType
}
