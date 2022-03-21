package teambuilder

import (
	"encoding/json"
	"net/http"

	"github.com/defry256/pokemon-helper/internal/app"
	"github.com/defry256/pokemon-helper/internal/errors"
	"github.com/defry256/pokemon-helper/internal/httpserver/handler"
	"github.com/defry256/pokemon-helper/internal/httpserver/response"
	"github.com/defry256/pokemon-helper/internal/pokemon"
	"github.com/defry256/pokemon-helper/internal/pokemontype"
)

const (
	MAX_TEAM_NUMBER = 6
)

func SimulateTeam(application *app.App) http.HandlerFunc {
	return handler.Handle(func(w http.ResponseWriter, r *http.Request) error {
		type payload struct {
			Pokemons           []string `json:"pokemons"`
			WithPokemonData    bool     `json:"with_pokemon_data"`
			WithTypeSuggestion bool     `json:"with_type_suggestion"`
		}

		var p *payload
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			return errors.NewBadRequestError(err.Error())
		}

		if len(p.Pokemons) > MAX_TEAM_NUMBER {
			return errors.NewBadRequestError("Team must not exceed 6 pokemons")
		}

		teamResponse := map[string]interface{}{}
		if p.WithPokemonData {
			pokemons := []*pokemon.PokemonData{}
			for _, p := range p.Pokemons {
				poke := application.Pokedex.GetPokedex(p)
				pokemons = append(pokemons, poke)
			}
			teamResponse["pokemons"] = pokemons
		}

		coveredTypes, uncoveredTypes, err := application.TeamBuilder.CalculateTypeCoverage(p.Pokemons)
		if err != nil {
			return err
		}
		teamResponse["covered_types"] = coveredTypes
		teamResponse["uncovered_types"] = uncoveredTypes

		if p.WithTypeSuggestion {
			suggestionCount := (MAX_TEAM_NUMBER - len(p.Pokemons)) * 2
			suggestionTypes := application.TeamBuilder.CalculateSuggestedType(uncoveredTypes, suggestionCount)
			teamResponse["suggestion_types"] = suggestionTypes
		}

		response.WithData(w, http.StatusOK, teamResponse)
		return nil
	})
}

func GetTypesSuggestion(application *app.App) http.HandlerFunc {
	return handler.Handle(func(w http.ResponseWriter, r *http.Request) error {
		type payload struct {
			UncoveredTypes   []string `json:"uncovered_types"`
			SuggestionLength int      `json:"suggestion_length"`
		}

		var p *payload
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			return errors.NewBadRequestError(err.Error())
		}

		uncoveredTypes := []pokemontype.IType{}
		for _, uncoveredType := range p.UncoveredTypes {
			uncoveredTypes = append(uncoveredTypes, pokemontype.Type(uncoveredType))
		}

		suggestionLength := 10
		if p.SuggestionLength != 0 {
			suggestionLength = p.SuggestionLength
		}

		suggestionTypes := application.TeamBuilder.CalculateSuggestedType(uncoveredTypes, suggestionLength)

		response.WithData(w, http.StatusOK, map[string]interface{}{
			"suggestion_types": suggestionTypes,
		})
		return nil
	})
}
