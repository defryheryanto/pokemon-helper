package pokemontype_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/defry256/pokemon-helper/internal/pokemontype"
)

func TestWeakness(test *testing.T) {
	types := []pokemontype.IType{
		pokemontype.FireType,
		pokemontype.FightingType,
		pokemontype.GhostType,
	}

	expectedWeakness := [][]pokemontype.IType{
		{
			pokemontype.WaterType,
			pokemontype.GroundType,
			pokemontype.RockType,
		},
		{
			pokemontype.FlyingType,
			pokemontype.PsychicType,
			pokemontype.FairyType,
		},
		{
			pokemontype.GhostType,
			pokemontype.DarkType,
		},
	}

	for i, t := range types {
		weakness := t.WeakAgainst()
		if !reflect.DeepEqual(weakness, expectedWeakness[i]) {
			test.Fatalf(fmt.Sprintf("weakness gotten not equal expected: %v -> %v != %v", t, weakness, expectedWeakness[i]))
		}
	}
}
