package pokedex

import (
	"net/http"

	"github.com/defry256/pokemon-helper/internal/app"
	"github.com/defry256/pokemon-helper/internal/errors"
	"github.com/defry256/pokemon-helper/internal/httpserver/handler"
	"github.com/defry256/pokemon-helper/internal/httpserver/response"
	"github.com/gorilla/mux"
)

func GetAllPokedex(application *app.App) http.HandlerFunc {
	return handler.Handle(func(w http.ResponseWriter, r *http.Request) error {
		search := r.URL.Query().Get("search")
		pokemons := application.Pokedex.GetAllPokedex(search)

		response.WithData(w, http.StatusOK, map[string]interface{}{
			"pokemons": pokemons,
		})
		return nil
	})
}

func GetPokedex(application *app.App) http.HandlerFunc {
	return handler.Handle(func(w http.ResponseWriter, r *http.Request) error {
		pokemoNname := mux.Vars(r)["pokemonName"]
		data := application.Pokedex.GetPokedex(pokemoNname)
		if data == nil {
			return errors.NewNotFoundError("Pokemon not found")
		}
		response.WithData(w, http.StatusOK, data)
		return nil
	})
}
