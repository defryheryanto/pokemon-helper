package pokedex

import (
	"net/http"
	"strconv"

	"github.com/defryheryanto/pokemon-helper/internal/app"
	"github.com/defryheryanto/pokemon-helper/internal/errors"
	"github.com/defryheryanto/pokemon-helper/internal/httpserver/handler"
	"github.com/defryheryanto/pokemon-helper/internal/httpserver/response"
	"github.com/defryheryanto/pokemon-helper/internal/pokedex"
	"github.com/gorilla/mux"
)

func GetAllPokedex(application *app.App) http.HandlerFunc {
	return handler.Handle(func(w http.ResponseWriter, r *http.Request) error {
		pageString := r.URL.Query().Get("page")
		page, err := strconv.Atoi(pageString)
		if err != nil || page == 0 {
			page = 1
		}

		pageSizeString := r.URL.Query().Get("pageSize")
		pageSize, err := strconv.Atoi(pageSizeString)
		if err != nil || pageSize == 0 {
			pageSize = 50
		}

		pokemons, err := application.Pokedex.GetAllPokedex(r.Context(), pokedex.GetAllPokedexFilter{
			Page:     page,
			PageSize: pageSize,
		})
		if err != nil {
			return err
		}

		response.WithData(w, http.StatusOK, map[string]interface{}{
			"pokemons": pokemons,
		})
		return nil
	})
}

func GetPokedex(application *app.App) http.HandlerFunc {
	return handler.Handle(func(w http.ResponseWriter, r *http.Request) error {
		pokemoNname := mux.Vars(r)["pokemonName"]
		data, err := application.Pokedex.GetPokedex(r.Context(), pokemoNname)
		if err != nil {
			return err
		}
		if data == nil {
			return errors.NewNotFoundError("Pokemon not found")
		}
		response.WithData(w, http.StatusOK, data)
		return nil
	})
}
