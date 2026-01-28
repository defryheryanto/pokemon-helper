package teambuilder

import (
	"context"
	"errors"
	"testing"

	appErrors "github.com/defryheryanto/pokemon-helper/internal/errors"
	pokedexmock "github.com/defryheryanto/pokemon-helper/internal/pokedex/mock"
	"github.com/defryheryanto/pokemon-helper/internal/pokemon"
	"github.com/defryheryanto/pokemon-helper/internal/pokemontype"
	"github.com/golang/mock/gomock"
)

type stubType struct {
	weak []pokemontype.IType
}

func (s stubType) WeakAgainst() []pokemontype.IType {
	return s.weak
}

func (s stubType) StrongAgainst() []pokemontype.IType {
	return nil
}

func TestCalculateTypeCoverageSuccess(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	pokedexService := pokedexmock.NewMockIService(ctrl)
	service := NewService(pokedexService)

	pokedexService.EXPECT().GetPokedex(gomock.Any(), "charmander").Return(&pokemon.PokemonData{
		ID:    4,
		Name:  "charmander",
		Types: []pokemontype.IType{pokemontype.FireType},
	}, nil)
	pokedexService.EXPECT().GetPokedex(gomock.Any(), "squirtle").Return(&pokemon.PokemonData{
		ID:    7,
		Name:  "squirtle",
		Types: []pokemontype.IType{pokemontype.WaterType},
	}, nil)

	covered, uncovered, err := service.CalculateTypeCoverage(context.Background(), []string{"charmander", "squirtle"})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	expectedCovered := append([]pokemontype.IType{}, pokemontype.FireType.StrongAgainst()...)
	expectedCovered = append(expectedCovered, pokemontype.WaterType.StrongAgainst()...)
	if len(covered) != len(expectedCovered) {
		t.Fatalf("expected covered length %d, got %d", len(expectedCovered), len(covered))
	}
	for i, value := range expectedCovered {
		if covered[i] != value {
			t.Fatalf("expected covered %v, got %v", expectedCovered, covered)
		}
	}

	expectedUncovered := typeSet(allTypes())
	for _, value := range expectedCovered {
		delete(expectedUncovered, value.(pokemontype.Type))
	}
	actualUncovered := typeSet(uncovered)
	if len(actualUncovered) != len(expectedUncovered) {
		t.Fatalf("expected uncovered length %d, got %d", len(expectedUncovered), len(actualUncovered))
	}
	for value := range expectedUncovered {
		if _, ok := actualUncovered[value]; !ok {
			t.Fatalf("expected uncovered to include %s", value)
		}
	}
}

func TestCalculateTypeCoverageNotFound(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	pokedexService := pokedexmock.NewMockIService(ctrl)
	service := NewService(pokedexService)

	pokedexService.EXPECT().GetPokedex(gomock.Any(), "missingno").Return(nil, nil)

	_, _, err := service.CalculateTypeCoverage(context.Background(), []string{"missingno"})
	var notFound appErrors.NotFoundError
	if err == nil || !errors.As(err, &notFound) {
		t.Fatalf("expected NotFoundError, got %v", err)
	}
}

func TestCalculateTypeCoveragePropagatesError(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	pokedexService := pokedexmock.NewMockIService(ctrl)
	service := NewService(pokedexService)

	pokeErr := errors.New("service failure")
	pokedexService.EXPECT().GetPokedex(gomock.Any(), "pikachu").Return(nil, pokeErr)

	_, _, err := service.CalculateTypeCoverage(context.Background(), []string{"pikachu"})
	if !errors.Is(err, pokeErr) {
		t.Fatalf("expected error %v, got %v", pokeErr, err)
	}
}

func TestCalculateSuggestedTypeDeterministic(t *testing.T) {
	t.Parallel()

	service := NewService(nil)

	uncovered := []pokemontype.IType{
		stubType{weak: []pokemontype.IType{pokemontype.FireType, pokemontype.FireType, pokemontype.FireType}},
		stubType{weak: []pokemontype.IType{pokemontype.WaterType, pokemontype.WaterType}},
	}

	result := service.CalculateSuggestedType(context.Background(), uncovered, 2)
	if len(result) != 2 {
		t.Fatalf("expected 2 suggested types, got %d", len(result))
	}
	if result[0] != pokemontype.FireType || result[1] != pokemontype.WaterType {
		t.Fatalf("expected [fire water], got %v", result)
	}
}

func allTypes() []pokemontype.IType {
	return []pokemontype.IType{
		pokemontype.NormalType,
		pokemontype.FireType,
		pokemontype.WaterType,
		pokemontype.ElectricType,
		pokemontype.GrassType,
		pokemontype.IceType,
		pokemontype.FightingType,
		pokemontype.PoisonType,
		pokemontype.GroundType,
		pokemontype.FlyingType,
		pokemontype.PsychicType,
		pokemontype.BugType,
		pokemontype.RockType,
		pokemontype.GhostType,
		pokemontype.DragonType,
		pokemontype.DarkType,
		pokemontype.SteelType,
		pokemontype.FairyType,
	}
}

func typeSet(types []pokemontype.IType) map[pokemontype.Type]struct{} {
	set := map[pokemontype.Type]struct{}{}
	for _, value := range types {
		if typed, ok := value.(pokemontype.Type); ok {
			set[typed] = struct{}{}
		}
	}
	return set
}
