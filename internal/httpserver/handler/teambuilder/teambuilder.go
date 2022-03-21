package teambuilder

import (
	"encoding/json"
	"net/http"

	"github.com/defry256/pokemon-helper/internal/app"
	"github.com/defry256/pokemon-helper/internal/errors"
	"github.com/defry256/pokemon-helper/internal/httpserver/handler"
	"github.com/defry256/pokemon-helper/internal/httpserver/response"
	"github.com/defry256/pokemon-helper/internal/pokemon"
)

func SimulateTeam(application *app.App) http.HandlerFunc {
	return handler.Handle(func(w http.ResponseWriter, r *http.Request) error {
		type payload struct {
			Pokemons []string `json:"pokemons"`
		}

		var p *payload
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			return errors.NewBadRequestError(err.Error())
		}

		if len(p.Pokemons) > 6 {
			return errors.NewBadRequestError("Team must not exceed 6 pokemons")
		}

		pokemons := []*pokemon.PokemonData{}
		for _, p := range p.Pokemons {
			poke := application.Pokedex.GetPokedex(p)
			pokemons = append(pokemons, poke)
		}

		coveredTypes, uncoveredTypes, err := application.TeamBuilder.CalculateTypeCoverage(p.Pokemons)
		if err != nil {
			return err
		}

		response.WithData(w, http.StatusOK, simulateTeamResponse{
			Pokemons:       pokemons,
			CoveredTypes:   coveredTypes,
			UncoveredTypes: uncoveredTypes,
		})
		return nil
	})
}
