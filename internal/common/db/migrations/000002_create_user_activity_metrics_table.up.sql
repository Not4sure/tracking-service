
CREATE TABLE user_activity_metrics (
  user_id BIGINT NOT NULL,
  event_count INTEGER NOT NULL,
  window_start_at TIMESTAMP NOT NULL,
  window_end_at TIMESTAMP NOT NULL,
  created_at TIMESTAMP NOT NULL,
  PRIMARY KEY (user_id, window_start_at)
);
