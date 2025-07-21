-- name: CreatePrimeCheck :execresult
INSERT INTO prime_checks (user_id, number_text, status) VALUES (?, ?, 'processing');

-- name: GetPrimeCheck :one
SELECT
    id,
    user_id,
    number_text,
    trace_id,
    message_id,
    is_prime,
    status,
    created_at,
    updated_at
FROM prime_checks
WHERE
    id = ?;

-- name: ListPrimeChecks :many
SELECT
    id,
    user_id,
    number_text,
    trace_id,
    message_id,
    is_prime,
    status,
    created_at,
    updated_at
FROM prime_checks
ORDER BY created_at DESC;

-- name: CreateOutboxMessage :execresult
INSERT INTO outbox (event_type, payload) VALUES (?, ?);

-- name: GetUnprocessedOutboxMessages :many
SELECT
    id,
    event_type,
    payload,
    processed,
    created_at,
    updated_at
FROM outbox
WHERE
    processed = FALSE
ORDER BY created_at ASC;

-- name: MarkOutboxMessageProcessed :exec
UPDATE outbox
SET
    processed = TRUE,
    updated_at = CURRENT_TIMESTAMP
WHERE
    id = ?;

-- name: UpdatePrimeCheckResult :exec
UPDATE prime_checks
SET
    trace_id = ?,
    message_id = ?,
    is_prime = ?,
    status = ?,
    updated_at = CURRENT_TIMESTAMP
WHERE
    id = ?;