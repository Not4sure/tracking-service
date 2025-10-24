package command

import (
	"context"

	"github.com/not4sure/tracking-service/internal/common/decorator"
	"github.com/not4sure/tracking-service/internal/domain/metric"
	"github.com/sirupsen/logrus"
)

type CalculateMetrics struct{}

type CalculateMetricsHandler decorator.CommandHandler[CalculateMetrics]

type calculateMetricsHandler struct {
	repo     metric.Repository
	provider metric.Provider
}

func NewCalculateMetricsHandler(
	repo metric.Repository,
	provider metric.Provider,
	logger *logrus.Entry,
	metricsClient decorator.MetricsClient,
) CalculateMetricsHandler {
	if repo == nil {
		panic("nil repo")
	}

	return decorator.ApplyCommandDecorators(
		calculateMetricsHandler{repo, provider},
		logger,
		metricsClient,
	)
}

func (h calculateMetricsHandler) Handle(ctx context.Context, cmd CalculateMetrics) error {
	tw := metric.PrevTimeWindow()

	metrics, err := h.provider.AtTimeWindow(ctx, tw)
	if err != nil {
		return err
	}

	for _, m := range metrics {
		err = h.repo.Store(ctx, m)
		if err != nil {
			return err
		}
	}

	return nil
}
