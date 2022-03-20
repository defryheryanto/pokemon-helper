package handler

import (
	"fmt"
	"net/http"

	"github.com/defry256/pokemon-helper/internal/httpserver/response"
	"github.com/defry256/pokemon-helper/internal/logger"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

func Handle(hf HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		err := hf(rw, r)
		if err != nil {
			//send error to analytics (if available)
			logger.Error(fmt.Sprintf("error handle - %v", err), err)
			response.WithError(rw, err)
		}
	}
}
