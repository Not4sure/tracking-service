package main

import (
	"context"

	"github.com/not4sure/tracking-service/internal/app/command"
	_ "github.com/not4sure/tracking-service/internal/common/logs"
	"github.com/not4sure/tracking-service/internal/service"
	"github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()
	app := service.NewApplication(ctx)

	logrus.Info("Running CalculateMetrics command")

	err := app.Commands.CalculateMetrics.Handle(ctx, command.CalculateMetrics{})
	if err != nil {
		logrus.WithError(err).Fatal("Cannot calculate user activity metrics")
	}
}
