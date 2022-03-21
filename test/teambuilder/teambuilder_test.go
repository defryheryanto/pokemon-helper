package teambuilder_test

import (
	"fmt"
	"testing"

	"github.com/defry256/pokemon-helper/internal/pokedex/v1"
	"github.com/defry256/pokemon-helper/internal/pokemontype"
	iteambuilder "github.com/defry256/pokemon-helper/internal/teambuilder"
	"github.com/defry256/pokemon-helper/internal/teambuilder/v1"
)

func setupService() iteambuilder.IService {
	pokedex := pokedex.NewService()
	service := teambuilder.NewService(pokedex)
	return service
}

func equalTypes(types1 []pokemontype.IType, types2 []pokemontype.IType, typesName string) error {
	if len(types1) != len(types2) {
		return fmt.Errorf("Length of %s is not match (%d != %d)", typesName, len(types1), len(types2))
	}

	for _, type1 := range types1 {
		notMatch := true
		for _, type2 := range types2 {
			if type1 == type2 {
				notMatch = false
			}
		}
		if notMatch {
			return fmt.Errorf("Type %s is not found on expected", type1)
		}
	}

	return nil
}

func TestTypeCoverage(test *testing.T) {
	pokemonNames := []string{"gengar", "machamp"}
	expectedCoveredTypes := []pokemontype.IType{
		pokemontype.PsychicType,
		pokemontype.GhostType,
		pokemontype.GrassType,
		pokemontype.FairyType,
		pokemontype.NormalType,
		pokemontype.IceType,
		pokemontype.RockType,
		pokemontype.DarkType,
		pokemontype.SteelType,
	}
	expectedUncoveredTypes := []pokemontype.IType{
		pokemontype.BugType,
		pokemontype.DragonType,
		pokemontype.ElectricType,
		pokemontype.FightingType,
		pokemontype.FireType,
		pokemontype.FlyingType,
		pokemontype.GroundType,
		pokemontype.PoisonType,
		pokemontype.WaterType,
	}

	service := setupService()
	coveredTypes, uncoveredTypes, err := service.CalculateTypeCoverage(pokemonNames)
	if err != nil {
		test.Fatal(err)
	}

	err = equalTypes(coveredTypes, expectedCoveredTypes, "coveredTypes")
	if err != nil {
		test.Fatal(err)
	}

	err = equalTypes(uncoveredTypes, expectedUncoveredTypes, "uncoveredTypes")
	if err != nil {
		test.Fatal(err)
	}
}

func TestSuggestedTypes(test *testing.T) {
	uncoveredTypes := []pokemontype.IType{
		pokemontype.NormalType,
		pokemontype.IceType,
		pokemontype.RockType,
		pokemontype.DarkType,
		pokemontype.SteelType,
		pokemontype.GroundType,
		pokemontype.FireType,
	}

	expectedSuggestedTypes := []pokemontype.IType{
		pokemontype.FightingType,
		pokemontype.GroundType,
		pokemontype.WaterType,
	}

	service := setupService()
	suggestedTypes := service.CalculateSuggestedType(uncoveredTypes, 3)
	err := equalTypes(suggestedTypes, expectedSuggestedTypes, "suggestedTypes")
	if err != nil {
		test.Fatal(err)
	}
}
