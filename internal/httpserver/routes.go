package httpserver

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/defry256/pokemon-helper/internal/app"
	"github.com/defry256/pokemon-helper/internal/httpserver/handler/pokedex"
	"github.com/defry256/pokemon-helper/internal/httpserver/handler/teambuilder"
	"github.com/gorilla/mux"
)

func HandleRoutes(a *app.App) http.Handler {
	root := mux.NewRouter()

	root.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		json.NewEncoder(rw).Encode("success")
	})

	root.HandleFunc("/api/v1/pokemon/{pokemonName}", pokedex.GetPokedex(a)).Methods(http.MethodGet)
	root.HandleFunc("/api/v1/simulate-team", teambuilder.SimulateTeam(a)).Methods(http.MethodGet)

	return http.TimeoutHandler(root, 30*time.Second, "Request Timeout")
}
