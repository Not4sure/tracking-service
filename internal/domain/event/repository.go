package event

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/not4sure/tracking-service/internal/common/errors"
)

var (
	ErrEventNotFound      = errors.NewNotFoundError("event not found", "Cannot found event")
	ErrEventAlreadyExists = errors.NewIncorrectInputError("event already exists", "Event with such id already exists")
)

// Repository is interface for storage of Events.
type Repository interface {
	Store(ctx context.Context, e *Event) error
	FindByUUID(ctx context.Context, uuid uuid.UUID) (*Event, error)
	List(ctx context.Context, userID uint, from time.Time, till time.Time) ([]*Event, error)
}
