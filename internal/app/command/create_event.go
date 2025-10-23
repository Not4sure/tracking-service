package command

import (
	"context"

	"github.com/not4sure/tracking-service/internal/common/decorator"
	"github.com/not4sure/tracking-service/internal/domain/event"
	"github.com/sirupsen/logrus"
)

type CreateEvent struct {
	UserID   uint
	Action   string
	Metadata map[string]string
}

type CreateEventHandler decorator.CommandHandler[CreateEvent]

type createEventHandler struct {
	repo event.Repository
}

func NewCreateEventHandler(
	repo event.Repository,
	logger *logrus.Entry,
	metricsClient decorator.MetricsClient,
) CreateEventHandler {
	if repo == nil {
		panic("nil repo")
	}

	return decorator.ApplyCommandDecorators(
		createEventHandler{repo},
		logger,
		metricsClient,
	)
}

func (h createEventHandler) Handle(ctx context.Context, cmd CreateEvent) error {
	e, err := event.New(cmd.UserID, cmd.Action, event.WithMetadata(cmd.Metadata))
	if err != nil {
		return err
	}

	return h.repo.Store(ctx, e)
}
