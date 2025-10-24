package http_api

import (
	"net/http"

	"github.com/not4sure/tracking-service/internal/app/command"
)

func (s APIServer) CalcuateMetrics(w http.ResponseWriter, r *http.Request) {
	cmd := command.CalculateMetrics{}
	err := s.app.Commands.CalculateMetrics.Handle(r.Context(), cmd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

