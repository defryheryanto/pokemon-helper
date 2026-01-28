package teambuilder

//go:generate go run github.com/golang/mock/mockgen -source=interface.go -package=mock -destination=mock/mock.go

import (
	"context"

	"github.com/defryheryanto/pokemon-helper/internal/pokemontype"
)

type IService interface {
	CalculateTypeCoverage(ctx context.Context, pokemonNames []string) (coveredTypes, uncoveredTypes []pokemontype.IType, err error)
	CalculateSuggestedType(ctx context.Context, uncoveredTypes []pokemontype.IType, suggestLength int) []pokemontype.IType
}
