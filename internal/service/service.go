package service

import (
	"context"

	"github.com/not4sure/tracking-service/internal/app"
)

func NewApplication(ctx context.Context) app.Application {
	// logger := logrus.NewEntry(logrus.StandardLogger())
	// metricsClient := metrics.NoOp{}

	return app.Application{
		Commands: app.Commands{},
		Queries:  app.Queries{},
	}
}
