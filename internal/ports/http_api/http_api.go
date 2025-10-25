package http_api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Masterminds/formenc/encoding/form"
	"github.com/go-chi/chi/v5"
	"github.com/not4sure/tracking-service/internal/app"
)

type APIServer struct {
	app app.Application
}

func NewAPIServer(app app.Application) APIServer {
	return APIServer{
		app: app,
	}
}

func (s APIServer) RegisterRoutes(r chi.Router) {
	r.Post("/event", s.CreateEvent)
	r.Get("/events", s.ListEvents)
	r.Post("/metrics", s.CalcuateMetrics)
	r.Get("/metrics", s.ListMetrics)
}

func unmarshallURLForm(r *http.Request, v any) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}

	return form.Unmarshal(r.Form, v)
}

func unmarshallBodyJSON(r *http.Request, v any) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, v)
}

func respondWithJSON(w http.ResponseWriter, statusCode int, data any) {
	sendHeaders(w, statusCode)
	json.NewEncoder(w).Encode(data)
}

func sendHeaders(w http.ResponseWriter, statusCode int) {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(statusCode)
}
