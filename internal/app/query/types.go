package query

import (
	"time"

	"github.com/google/uuid"
	"github.com/not4sure/tracking-service/internal/domain/event"
	"github.com/not4sure/tracking-service/internal/domain/metric"
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

type Metric struct {
	UserID        uint
	EventCount    uint
	WindowStartAt time.Time
	WindowEndAt   time.Time
	CreatedAt     time.Time
}

func domainMetricToView(m *metric.Metric) Metric {
	return Metric{
		UserID:        m.UserID(),
		EventCount:    m.EventCount(),
		WindowStartAt: m.TimeWindow().Start(),
		WindowEndAt:   m.TimeWindow().End(),
		CreatedAt:     m.CreatedAt(),
	}
}
