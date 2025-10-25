package adapters

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"time"

	"github.com/emicklei/pgtalk/convert"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/not4sure/tracking-service/internal/common/db"
	"github.com/not4sure/tracking-service/internal/domain/event"
)

type EventsPostgresRepository struct {
	conn    *pgxpool.Pool
	queries *db.Queries
}

func NewEventsPostgresRepository(conn *pgxpool.Pool) *EventsPostgresRepository {
	if conn == nil {
		panic("nil db")
	}

	return &EventsPostgresRepository{
		conn:    conn,
		queries: db.New(conn),
	}
}

func (pr *EventsPostgresRepository) Store(ctx context.Context, e *event.Event) error {
	metadata, err := json.Marshal(e.Metadata())
	if err != nil {
		return err
	}

	arg := db.CreateEventParams{
		ID:        convert.UUID(e.UUID()),
		OccuredAt: convert.TimeToTimestamp(e.OccuredAt()),
		UserID:    int64(e.UserID()),
		Action:    e.Action(),
		Metadata:  metadata,
	}

	err = pr.queries.CreateEvent(ctx, arg)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case "23505": // unique violation
				return event.ErrEventAlreadyExists
			}
		}

	}

	return nil
}

func (pr *EventsPostgresRepository) List(ctx context.Context, userID uint, from, till time.Time) ([]*event.Event, error) {
	arg := db.ListEventsParams{
		UserID:      int64(userID),
		OccuredAt:   convert.TimeToTimestamp(from),
		OccuredAt_2: convert.TimeToTimestamp(till),
	}

	ee, err := pr.queries.ListEvents(ctx, arg)
	if err != nil {
		return nil, err
	}

	domainEvents := []*event.Event{}
	for _, e := range ee {
		domainEvent, err := pr.UnmarshalEvent(e)
		if err != nil {
			return nil, err
		}
		domainEvents = append(domainEvents, domainEvent)
	}

	return domainEvents, nil
}

func (pr *EventsPostgresRepository) FindByUUID(ctx context.Context, id uuid.UUID) (*event.Event, error) {
	e, err := pr.queries.FindByID(ctx, convert.UUID(id))
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, event.ErrEventNotFound
		}

		return nil, err
	}

	return pr.UnmarshalEvent(e)
}

func (pr EventsPostgresRepository) UnmarshalEvent(mdl db.Event) (*event.Event, error) {
	metadata := map[string]string{}
	err := json.Unmarshal(mdl.Metadata, &metadata)
	if err != nil {
		return nil, err
	}

	return event.UnmarshalEventFromDatabase(
		mdl.ID.Bytes,
		mdl.OccuredAt.Time,
		uint(mdl.UserID),
		mdl.Action,
		metadata,
	)
}

func NewPostgresConnection(ctx context.Context) (*pgxpool.Pool, error) {
	url := os.Getenv("POSTGRES_URL")

	return pgxpool.New(ctx, url)
}
