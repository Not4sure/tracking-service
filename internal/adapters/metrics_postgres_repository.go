package adapters

import (
	"context"
	"time"

	"github.com/emicklei/pgtalk/convert"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/not4sure/tracking-service/internal/common/db"
	"github.com/not4sure/tracking-service/internal/domain/metric"
)

type MetricsPostgresRepository struct {
	conn *pgxpool.Pool
}

func NewMetricsPostgresRepository(conn *pgxpool.Pool) MetricsPostgresRepository {
	return MetricsPostgresRepository{
		conn,
	}
}

func (r *MetricsPostgresRepository) Store(ctx context.Context, m *metric.Metric) error {
	arg := db.UpsertUserActivityMetricParams{
		UserID:        int64(m.UserID()),
		EventCount:    int32(m.EventCount()),
		WindowStartAt: convert.TimeToTimestamp(m.TimeWindow().Start()),
		WindowEndAt:   convert.TimeToTimestamp(m.TimeWindow().End()),
		CreatedAt:     convert.TimeToTimestamp(m.CreatedAt()),
	}

	return r.queries().UpsertUserActivityMetric(ctx, arg)
}

func (r *MetricsPostgresRepository) List(ctx context.Context, userID uint, from, till time.Time) ([]*metric.Metric, error) {
	arg := db.ListUserActivityMetricsParams{
		UserID:          int64(userID),
		WindowStartAt:   convert.TimeToTimestamp(from),
		WindowStartAt_2: convert.TimeToTimestamp(till),
	}

	mm, err := r.queries().ListUserActivityMetrics(ctx, arg)
	if err != nil {
		return nil, err
	}

	domainMetrics := []*metric.Metric{}
	for _, m := range mm {
		domainMetrics = append(domainMetrics, r.UnmarshalMetric(m))
	}

	return domainMetrics, nil
}

func (r *MetricsPostgresRepository) queries() *db.Queries {
	return db.New(r.conn)
}

func (r MetricsPostgresRepository) UnmarshalMetric(m db.UserActivityMetric) *metric.Metric {
	return metric.UnmarshalMetricFromDatabase(
		uint(m.UserID),
		uint(m.EventCount),
		m.WindowStartAt.Time,
		m.CreatedAt.Time,
	)

}
