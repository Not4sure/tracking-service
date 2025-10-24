
-- name: CreateEvent :exec
INSERT INTO events (
  id, occured_at, user_id, action, metadata
) VALUES (
  $1, $2, $3, $4, $5
);

-- name: ListEvents :many
SELECT * 
FROM events 
WHERE user_id = $1 
  AND occured_at BETWEEN $2 AND $3
ORDER BY occured_at;

-- name: FindByID :one
SELECT *
FROM events
WHERE id = $1 
LIMIT 1;
