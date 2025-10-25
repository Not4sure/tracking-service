package http_api

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/not4sure/tracking-service/internal/app/query"
	"github.com/not4sure/tracking-service/internal/server/httperr"
)

type listMetricsParams struct {
	UserID uint      `form:"user_id"`
	From   time.Time `form:"from"`
	Till   time.Time `form:"till"`
}

func (p *listMetricsParams) FormSetFrom(v []string) error {
	if len(v) < 1 {
		return errors.New("no from field")
	}

	t, err := time.Parse(time.RFC3339, v[0])
	p.From = t

	return err
}

func (p *listMetricsParams) FormSetTill(v []string) error {
	fmt.Println(len(v))
	if len(v) < 1 {
		return errors.New("no till field")
	}

	t, err := time.Parse(time.RFC3339, v[0])
	p.Till = t

	return err
}

func (p listMetricsParams) toQuery() query.ListMetrics {
	return query.ListMetrics{
		UserID: p.UserID,
		From:   p.From,
		Till:   p.Till,
	}
}

func (s *APIServer) ListMetrics(w http.ResponseWriter, r *http.Request) {
	params := listMetricsParams{}
	if err := unmarshallURLForm(r, &params); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mm, err := s.app.Queries.ListMetrics.Handle(r.Context(), params.toQuery())
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	respondWithJSON(w, http.StatusOK, appMetricsToResponse(mm))
}

func appMetricsToResponse(mm []query.Metric) Metrics {
	rsp := Metrics{Metrics: []Metric{}}

	for _, m := range mm {
		rsp.Metrics = append(rsp.Metrics, Metric{
			UserID:        m.UserID,
			EventCount:    m.EventCount,
			WindowStartAt: m.WindowStartAt,
			WindowEndAt:   m.WindowEndAt,
			CreatedAt:     m.CreatedAt,
		})
	}

	return rsp
}
