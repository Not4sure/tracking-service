package query

import (
	"time"

	"github.com/google/uuid"
	"github.com/not4sure/tracking-service/internal/domain/event"
)

type Event struct {
	UUID      uuid.UUID
	OccuredAt time.Time

	UserID   uint
	Action   string
	Metadata map[string]string
}

func domainEventToView(e *event.Event) Event {
	return Event{
		UUID:      e.UUID(),
		OccuredAt: e.OccuredAt(),

		UserID:   e.UserID(),
		Action:   e.Action(),
		Metadata: e.Metadata(),
	}
}
