-- name: AddSession :exec
INSERT INTO "sessions"("token", "player_id")
  VALUES($1, $2);
