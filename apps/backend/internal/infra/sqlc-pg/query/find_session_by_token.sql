-- name: FindSessionByToken :one
SELECT sqlc.embed(sessions) FROM "sessions" WHERE "token" = $1;
