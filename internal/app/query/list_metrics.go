package query

import (
	"context"
	"time"

	"github.com/not4sure/tracking-service/internal/common/decorator"
	"github.com/not4sure/tracking-service/internal/domain/metric"
	"github.com/sirupsen/logrus"
)

type ListMetrics struct {
	UserID uint
	From   time.Time
	Till   time.Time
}

type ListMetricsHandler decorator.QueryHandler[ListMetrics, []Metric]

type listMetricsHandler struct {
	repo metric.Repository
}

func NewListMetricsHandler(
	repo metric.Repository,
	logger *logrus.Entry,
	metricsClient decorator.MetricsClient,
) ListMetricsHandler {
	if repo == nil {
		panic("nil repo")
	}

	return decorator.ApplyQueryDecorators(
		listMetricsHandler{repo},
		logger,
		metricsClient,
	)
}

func (h listMetricsHandler) Handle(ctx context.Context, query ListMetrics) ([]Metric, error) {
	domainMetrics, err := h.repo.List(ctx, query.UserID, query.From, query.Till)
	if err != nil {
		return nil, err
	}

	mm := []Metric{}
	for _, m := range domainMetrics {
		mm = append(mm, domainMetricToView(m))
	}

	return mm, nil
}
