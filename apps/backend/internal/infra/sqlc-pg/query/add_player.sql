-- name: AddPlayer :exec
INSERT INTO "players"("display_name", "pass_hash", "login")
  VALUES($1, $2, $3) RETURNING ("id");
