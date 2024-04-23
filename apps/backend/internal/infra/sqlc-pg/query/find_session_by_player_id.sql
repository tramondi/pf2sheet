-- name: FindSessionByPlayerID :one
SELECT sqlc.embed(sessions) FROM "sessions" WHERE "player_id" = $1;
