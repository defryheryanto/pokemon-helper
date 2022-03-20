package pokedex

import (
	"net/http"

	"github.com/defry256/pokemon-helper/internal/app"
	"github.com/defry256/pokemon-helper/internal/httpserver/handler"
	"github.com/defry256/pokemon-helper/internal/httpserver/response"
	"github.com/gorilla/mux"
)

func GetPokedex(application *app.App) http.HandlerFunc {
	return handler.Handle(func(w http.ResponseWriter, r *http.Request) error {
		pokemoNname := mux.Vars(r)["pokemonName"]
		data := application.Pokedex.GetPokedex(pokemoNname)
		response.WithData(w, http.StatusOK, data)
		return nil
	})
}
