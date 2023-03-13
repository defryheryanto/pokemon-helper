package httpserver

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/defry256/pokemon-helper/internal/app"
	"github.com/defry256/pokemon-helper/internal/httpserver/handler/pokedex"
	"github.com/defry256/pokemon-helper/internal/httpserver/handler/teambuilder"
	"github.com/defry256/pokemon-helper/internal/tracer"
	"github.com/gorilla/mux"
	"go.opentelemetry.io/otel/trace"
)

func HandleRoutes(a *app.App, tr trace.Tracer) http.Handler {
	root := mux.NewRouter()

	root.Use(tracer.TracerMiddleware(tr))
	root.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		json.NewEncoder(rw).Encode("success")
	})

	root.HandleFunc("/api/v1/pokemons", pokedex.GetAllPokedex(a)).Methods(http.MethodGet)
	root.HandleFunc("/api/v1/pokemons/{pokemonName}", pokedex.GetPokedex(a)).Methods(http.MethodGet)
	root.HandleFunc("/api/v1/simulate-team", teambuilder.SimulateTeam(a)).Methods(http.MethodGet)
	root.HandleFunc("/api/v1/type/suggestion", teambuilder.GetTypesSuggestion(a)).Methods(http.MethodGet)

	return http.TimeoutHandler(root, 30*time.Second, "Request Timeout")
}
