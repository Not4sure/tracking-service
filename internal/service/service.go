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
	conn, err := adapters.NewPostgresConnection(ctx)
	if err != nil {
		panic(err)
	}

	eventRepo := adapters.NewEventsPostgresRepository(conn)
	metricsRepo := adapters.NewMetricsPostgresRepository(conn)
	metricsProvider := adapters.NewMetricsPostgresProvider(conn)

	logger := logrus.NewEntry(logrus.StandardLogger())
	appMetricsClient := metrics.NoOp{}

	return app.Application{
		Commands: app.Commands{
			CreateEvent:      command.NewCreateEventHandler(eventRepo, logger, appMetricsClient),
			CalculateMetrics: command.NewCalculateMetricsHandler(metricsRepo, metricsProvider, logger, appMetricsClient),
		},
		Queries: app.Queries{
			ListEvents:  query.NewListEventsHandler(eventRepo, logger, appMetricsClient),
			ListMetrics: query.NewListMetricsHandler(metricsRepo, logger, appMetricsClient),
		},
	}
}
