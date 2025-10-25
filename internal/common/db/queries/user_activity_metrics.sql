
-- name: ListUserActivityMetrics :many
SELECT * 
FROM user_activity_metrics 
WHERE user_id = $1
  AND window_start_at BETWEEN $2 AND $3
ORDER BY window_start_at;

-- name: UpsertUserActivityMetric :exec
INSERT INTO user_activity_metrics (
  user_id, event_count, window_start_at, window_end_at, created_at
) VALUES (
  $1, $2, $3, $4, $5
) ON CONFLICT (user_id, window_start_at) DO UPDATE
SET 
  event_count = EXCLUDED.event_count,
  created_at = EXCLUDED.created_at;

