package service

import (
	"context"

	"github.com/not4sure/tracking-service/internal/adapters"
	"github.com/not4sure/tracking-service/internal/app"
	"github.com/not4sure/tracking-service/internal/app/command"
	"github.com/not4sure/tracking-service/internal/app/query"
	"github.com/not4sure/tracking-service/internal/common/metrics"
	"github.com/sirupsen/logrus"
)

func NewApplication(ctx context.Context) app.Application {
	eventRepo := adapters.NewEventMemoryRepository()
	logger := logrus.NewEntry(logrus.StandardLogger())
	metricsClient := metrics.NoOp{}

	return app.Application{
		Commands: app.Commands{
			CreateEvent: command.NewCreateEventHandler(eventRepo, logger, metricsClient),
		},
		Queries: app.Queries{
			ListEvents: query.NewListEventsHandler(eventRepo, logger, metricsClient),
		},
	}
}
