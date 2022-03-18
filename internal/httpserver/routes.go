package httpserver

import (
	"encoding/json"
	"net/http"

	"github.com/defry256/pokemon-helper/internal/app"
	"github.com/gorilla/mux"
)

func HandleRoutes(a *app.App) http.Handler {
	root := mux.NewRouter()

	root.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		json.NewEncoder(rw).Encode("success")
	})

	return root
}
