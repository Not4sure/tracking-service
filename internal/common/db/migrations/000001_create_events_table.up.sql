
CREATE TABLE events (
  id UUID PRIMARY KEY,
  occured_at TIMESTAMP NOT NULL,
  user_id BIGINT NOT NULL,
  action TEXT NOT NULL,
  metadata JSONB
);

CREATE INDEX idx_events_user_id_occured_at 
  ON events (user_id, occured_at);
