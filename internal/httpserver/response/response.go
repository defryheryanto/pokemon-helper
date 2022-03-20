package response

import (
	"encoding/json"
	"net/http"

	"github.com/defry256/pokemon-helper/internal/errors"
)

func WithData(rw http.ResponseWriter, responseStatus int, data interface{}) {
	(rw).Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.WriteHeader(responseStatus)
	json.NewEncoder(rw).Encode(data)
}

func WithError(rw http.ResponseWriter, err error) {
	(rw).Header().Set("Content-Type", "application/json; charset=utf-8")
	code, msg := getError(err)
	rw.WriteHeader(code)
	json.NewEncoder(rw).Encode(map[string]string{
		"error": msg,
	})
}

func getError(err error) (int, string) {
	switch e := err.(type) {
	case errors.UnauthorizedError:
		return http.StatusUnauthorized, e.Message()
	case errors.BadRequestError:
		return http.StatusBadRequest, e.Message()
	case errors.ForbiddenError:
		return http.StatusForbidden, e.Message()
	case errors.NotFoundError:
		return http.StatusNotFound, e.Message()
	default:
		return http.StatusInternalServerError, e.Error()
	}
}
