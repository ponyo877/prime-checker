-- name: CreatePrimeCheck :execresult
INSERT INTO prime_requests (user_id, number_text)
VALUES (?, ?);

-- name: GetPrimeCheck :one
SELECT id, user_id, number_text, created_at, updated_at
FROM prime_requests
WHERE id = ?;

-- name: ListPrimeChecks :many
SELECT id, user_id, number_text, created_at, updated_at
FROM prime_requests
ORDER BY created_at DESC;

-- name: CreateOutboxMessage :execresult
INSERT INTO outbox (event_type, payload)
VALUES (?, ?);

-- name: GetUnprocessedOutboxMessages :many
SELECT id, event_type, payload, processed, created_at, updated_at
FROM outbox
WHERE processed = FALSE
ORDER BY created_at ASC;

-- name: MarkOutboxMessageProcessed :exec
UPDATE outbox
SET processed = TRUE, updated_at = CURRENT_TIMESTAMP
WHERE id = ?;