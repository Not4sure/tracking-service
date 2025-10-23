package event

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrEventNotFound      = errors.New("event not found")
	ErrEventAlreadyExists = errors.New("event already exists")
)

// Repository is interface for storage of Events.
type Repository interface {
	Store(ctx context.Context, e *Event) error
	FindByUUID(ctx context.Context, uuid uuid.UUID) (*Event, error)
	List(ctx context.Context, userID uint, from time.Time, till time.Time) ([]*Event, error)
}
