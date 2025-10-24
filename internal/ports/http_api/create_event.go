package http_api

import (
	"net/http"

	"github.com/not4sure/tracking-service/internal/app/command"
)

type createEventParams struct {
	UserID   uint              `json:"user_id"`
	Action   string            `json:"action"`
	Metadata map[string]string `json:"metadata"`
}

func (p createEventParams) toCmd() command.CreateEvent {
	return command.CreateEvent{
		UserID:   p.UserID,
		Action:   p.Action,
		Metadata: p.Metadata,
	}
}

func (s APIServer) CreateEvent(w http.ResponseWriter, r *http.Request) {
	params := createEventParams{}
	if err := unmarshallBodyJSON(r, &params); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := s.app.Commands.CreateEvent.Handle(r.Context(), params.toCmd())
	if err != nil {
		// TODO: handle application error.
		return
	}

	respondWithJSON(w, http.StatusCreated, struct{ Status string }{Status: "OK"})
}
