package httpserver

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/defryheryanto/pokemon-helper/internal/app"
	"github.com/defryheryanto/pokemon-helper/internal/httpserver/handler/pokedex"
	"github.com/defryheryanto/pokemon-helper/internal/httpserver/handler/teambuilder"
	"github.com/defryheryanto/pokemon-helper/internal/tracer"
	"github.com/gorilla/mux"
	"go.opentelemetry.io/otel/trace"
)

func HandleRoutes(a *app.App, tr trace.Tracer) http.Handler {
	root := mux.NewRouter()

	root.Use(tracer.TracerMiddleware(tr))
	root.Use(corsMiddleware)
	root.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		json.NewEncoder(rw).Encode("success")
	})

	root.HandleFunc("/api/v1/pokemons", pokedex.GetAllPokedex(a)).Methods(http.MethodGet)
	root.HandleFunc("/api/v1/pokemons/{pokemonName}", pokedex.GetPokedex(a)).Methods(http.MethodGet)
	root.HandleFunc("/api/v1/simulate-team", teambuilder.SimulateTeam(a)).Methods(http.MethodPost)
	root.HandleFunc("/api/v1/type/suggestion", teambuilder.GetTypesSuggestion(a)).Methods(http.MethodGet)

	return http.TimeoutHandler(root, 30*time.Second, "Request Timeout")
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
