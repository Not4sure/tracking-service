package query

import (
	"context"
	"time"

	"github.com/not4sure/tracking-service/internal/common/decorator"
	"github.com/not4sure/tracking-service/internal/domain/event"
	"github.com/sirupsen/logrus"
)

type ListEvents struct {
	UserID uint
	From   time.Time
	Till   time.Time
}

type ListEventsHandler decorator.QueryHandler[ListEvents, []Event]

type listEventsHandler struct {
	repo event.Repository
}

func NewListEventsHandler(
	repo event.Repository,
	logger *logrus.Entry,
	metricsClient decorator.MetricsClient,
) ListEventsHandler {
	return decorator.ApplyQueryDecorators(
		listEventsHandler{repo},
		logger,
		metricsClient,
	)
}

func (h listEventsHandler) Handle(ctx context.Context, query ListEvents) ([]Event, error) {
	domainEvents, err := h.repo.List(ctx, query.UserID, query.From, query.Till)
	if err != nil {
		return []Event{}, err
	}

	ee := []Event{}
	for _, e := range domainEvents {
		ee = append(ee, domainEventToView(e))
	}

	return ee, nil
}

