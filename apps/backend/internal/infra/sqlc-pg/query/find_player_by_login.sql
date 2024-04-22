-- name: FindPlayerByLogin :one
SELECT sqlc.embed(players) FROM "players" WHERE "login" = $1;
