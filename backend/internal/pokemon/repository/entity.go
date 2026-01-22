package repository

import (
	"encoding/json"
	"fmt"

	"github.com/defryheryanto/pokemon-helper/internal/pokemon"
	"github.com/defryheryanto/pokemon-helper/internal/pokemontype"
)

type PokemonEntity struct {
	ID         int
	Name       string
	BaseStatus []byte
	Types      []byte
	Sprites    string
}

func newPokemonEntityFromModel(model *pokemon.PokemonData) (*PokemonEntity, error) {
	if model == nil {
		return nil, nil
	}

	var baseStatus []byte
	if model.BaseStatus != nil {
		data, err := json.Marshal(model.BaseStatus)
		if err != nil {
			return nil, err
		}
		baseStatus = data
	}

	rawTypes := make([]string, 0, len(model.Types))
	for _, t := range model.Types {
		rawTypes = append(rawTypes, string(pokemontype.Type(fmt.Sprintf("%v", t))))
	}

	var types []byte
	if len(rawTypes) > 0 {
		data, err := json.Marshal(rawTypes)
		if err != nil {
			return nil, err
		}
		types = data
	}

	return &PokemonEntity{
		ID:         model.ID,
		Name:       model.Name,
		BaseStatus: baseStatus,
		Types:      types,
		Sprites:    model.Sprites,
	}, nil
}

func (e *PokemonEntity) ToModel() (*pokemon.PokemonData, error) {
	if e == nil {
		return nil, nil
	}

	var baseStatus *pokemon.Status
	if len(e.BaseStatus) > 0 {
		var status pokemon.Status
		if err := json.Unmarshal(e.BaseStatus, &status); err != nil {
			return nil, err
		}
		baseStatus = &status
	}

	pokemonTypes := []pokemontype.IType{}
	if len(e.Types) > 0 {
		var rawTypes []string
		if err := json.Unmarshal(e.Types, &rawTypes); err != nil {
			return nil, err
		}
		for _, t := range rawTypes {
			pokemonTypes = append(pokemonTypes, pokemontype.Type(t))
		}
	}

	return &pokemon.PokemonData{
		ID:         e.ID,
		Name:       e.Name,
		BaseStatus: baseStatus,
		Types:      pokemonTypes,
		Sprites:    e.Sprites,
	}, nil
}
