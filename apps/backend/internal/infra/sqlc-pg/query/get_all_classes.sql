-- name: GetAllClasses :many
SELECT sqlc.embed(classes) FROM "classes";
