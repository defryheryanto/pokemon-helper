package pokemontype_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/defry256/pokemon-helper/internal/pokemontype"
)

func TestEffective(test *testing.T) {
	types := []pokemontype.IType{
		pokemontype.FireType,
		pokemontype.FightingType,
		pokemontype.GhostType,
	}

	expectedEffective := [][]pokemontype.IType{
		{
			pokemontype.GrassType,
			pokemontype.IceType,
			pokemontype.BugType,
			pokemontype.SteelType,
		},
		{
			pokemontype.NormalType,
			pokemontype.IceType,
			pokemontype.RockType,
			pokemontype.DarkType,
			pokemontype.SteelType,
		},
		{
			pokemontype.PsychicType,
			pokemontype.GhostType,
		},
	}

	for i, t := range types {
		strength := t.StrongAgainst()
		if !reflect.DeepEqual(strength, expectedEffective[i]) {
			test.Fatalf(fmt.Sprintf("weakness gotten not equal expected: %v -> %v != %v", t, strength, expectedEffective[i]))
		}
	}
}
