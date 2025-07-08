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