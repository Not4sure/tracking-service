package http_api

import (
	"errors"
	"net/http"
	"time"

	"github.com/not4sure/tracking-service/internal/app/query"
)

type listEventsParams struct {
	UserID uint      `form:"user_id"`
	From   time.Time `form:"from"`
	Till   time.Time `form:"till"`
}

func (p *listEventsParams) FormSetFrom(v []string) error {
	if len(v) < 1 {
		return errors.New("no from field")
	}

	from, err := time.Parse(time.RFC3339, v[0])
	p.From = from

	return err
}

func (p *listEventsParams) FormSetTill(v []string) error {
	if len(v) < 1 {
		return errors.New("no till field")
	}

	date, err := time.Parse(time.RFC3339, v[0])
	p.Till = date

	return err
}

func (p listEventsParams) toQuery() query.ListEvents {
	return query.ListEvents{
		UserID: p.UserID,
		From:   p.From,
		Till:   p.Till,
	}
}

func (s APIServer) ListEvents(w http.ResponseWriter, r *http.Request) {
	params := listEventsParams{}
	if err := unmarshallURLForm(r, &params); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ee, err := s.app.Queries.ListEvents.Handle(r.Context(), params.toQuery())
	if err != nil {
		// TODO: handle application error.
		return
	}

	respondWithJSON(w, http.StatusOK, appEventsToResponse(ee))
}

func appEventsToResponse(ee []query.Event) Events {
	rsp := Events{}
	for _, e := range ee {
		rsp.Events = append(rsp.Events, Event{
			UUID:      e.UUID.String(),
			OccuredAt: e.OccuredAt,
			UserID:    e.UserID,
			Action:    e.Action,
			Metadata:  e.Metadata,
		})
	}

	return rsp
}
