package httperr

import (
	"encoding/json"
	"net/http"

	"github.com/not4sure/tracking-service/internal/common/errors"
)

func InternalError(slug string, err error, w http.ResponseWriter, r *http.Request) {
	httpRespondWithError(err, slug, w, r, "Internal server error", http.StatusInternalServerError)
}

func Unauthorised(slug string, err error, w http.ResponseWriter, r *http.Request) {
	httpRespondWithError(err, slug, w, r, "Unauthorised", http.StatusUnauthorized)
}

func NotFound(slug string, err error, w http.ResponseWriter, r *http.Request) {
	httpRespondWithError(err, slug, w, r, "Not Found", http.StatusNotFound)
}

func BadRequest(slug string, err error, w http.ResponseWriter, r *http.Request) {
	httpRespondWithError(err, slug, w, r, "Bad request", http.StatusBadRequest)
}

func RespondWithSlugError(err error, w http.ResponseWriter, r *http.Request) {
	slugError, ok := err.(errors.SlugError)
	if !ok {
		InternalError("internal-server-error", err, w, r)
		return
	}

	switch slugError.ErrorType() {
	case errors.ErrorTypeAuthorization:
		Unauthorised(slugError.Slug(), slugError, w, r)
	case errors.ErrorTypeNotFound:
		NotFound(slugError.Slug(), slugError, w, r)
	case errors.ErrorTypeIncorrectInput:
		BadRequest(slugError.Slug(), slugError, w, r)
	default:
		InternalError(slugError.Slug(), slugError, w, r)
	}
}

func httpRespondWithError(err error, slug string, w http.ResponseWriter, r *http.Request, logMSg string, status int) {
	resp := ErrorResponse{slug, status}

	err = resp.Render(w, r)
	if err != nil {
		panic(err)
	}
}

type ErrorResponse struct {
	Slug       string `json:"slug"`
	httpStatus int
}

func (e ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(e.httpStatus)

	bb, err := json.Marshal(e)
	w.Write(bb)
	return err
}
