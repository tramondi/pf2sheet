-- name: DeleteSessionByToken :exec
DELETE FROM "sessions" WHERE "token" = $1;
