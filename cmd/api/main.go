package main

import (
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/go-co-op/gocron/v2"
	"github.com/not4sure/tracking-service/internal/app"
	"github.com/not4sure/tracking-service/internal/app/command"
	"github.com/not4sure/tracking-service/internal/ports/http_api"
	"github.com/not4sure/tracking-service/internal/server"
	"github.com/not4sure/tracking-service/internal/service"
	"github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()

	app := service.NewApplication(ctx)

	shotdown, err := startCalculateMetricsJob(app)
	if err != nil {
		logrus.Fatal(err)
	}
	defer shotdown()

	server.RunServer(func(router chi.Router) {
		apiServer := http_api.NewAPIServer(app)
		apiServer.RegisterRoutes(router)
	})
}

func startCalculateMetricsJob(app app.Application) (func(), error) {
	s, err := gocron.NewScheduler()
	if err != nil {
		return nil, err
	}

	_, err = s.NewJob(
		gocron.CronJob("1 */4 * * *", false), // at 1 minute every 4th hour
		gocron.NewTask(func() {
			app.Commands.CalculateMetrics.Handle(context.Background(), command.CalculateMetrics{})
		}),
	)
	if err != nil {
		return nil, err
	}

	s.Start()

	return func() {
		s.Shutdown()
	}, nil
}
