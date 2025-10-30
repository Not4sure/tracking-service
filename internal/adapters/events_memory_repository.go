package adapters

import (
	"context"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/not4sure/tracking-service/internal/domain/event"
)

type EventsMemoryRepository struct {
	sync.Mutex
	events map[uuid.UUID]event.Event
}

func NewEventsMemoryRepository() event.Repository {
	return &EventsMemoryRepository{
		events: map[uuid.UUID]event.Event{},
	}
}

func (mr *EventsMemoryRepository) Store(_ context.Context, e *event.Event) error {
	mr.Lock()
	defer mr.Unlock()

	if _, ok := mr.events[e.UUID()]; ok {
		return event.ErrEventAlreadyExists
	}

	mr.events[e.UUID()] = *e
	return nil
}

func (mr *EventsMemoryRepository) FindByUUID(_ context.Context, id uuid.UUID) (*event.Event, error) {
	mr.Lock()
	defer mr.Unlock()

	e, ok := mr.events[id]
	if !ok {
		return nil, event.ErrEventNotFound
	}

	return &e, nil
}

func (mr *EventsMemoryRepository) List(_ context.Context, userID uint, from time.Time, till time.Time) ([]*event.Event, error) {
	var ee []*event.Event
	for _, e := range mr.events {
		switch true {
		case e.UserID() != userID:
			continue
		case e.OccuredAt().Before(from):
			continue
		case e.OccuredAt().After(till):
			continue
		}

		ee = append(ee, &e)
	}

	return ee, nil
}
