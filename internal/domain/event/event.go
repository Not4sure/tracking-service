package event

import (
	"time"

	"github.com/google/uuid"
)

// Event
type Event struct {
	uuid      uuid.UUID
	occuredAt time.Time
	userID    uint
	action    Action

	metadata Metadata
}

func New(userID uint, action string, opts ...EventOption) (*Event, error) {
	e := &Event{
		uuid:      uuid.New(),
		occuredAt: time.Now(),
		userID:    userID,
		action:    Action{action},
		metadata:  map[string]string{},
	}

	for _, opt := range opts {
		opt(e)
	}

	return e, nil
}

// UUID returns id of Event.
func (e *Event) UUID() uuid.UUID {
	return e.uuid
}

// OccuredAt returns time at which event had occured.
func (e *Event) OccuredAt() time.Time {
	return e.occuredAt
}

// UserID is id of a user which triggered Event.
func (e *Event) UserID() uint {
	return e.userID
}

// Action is a string representation of action Event is associated with.
func (e *Event) Action() string {
	return e.action.String()
}

// Metadata is a map of metadata values associated with Event.
func (e *Event) Metadata() Metadata {
	return e.metadata
}
