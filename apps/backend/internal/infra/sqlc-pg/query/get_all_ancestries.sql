-- name: GetAllAncestries :many
SELECT sqlc.embed(ancestries) FROM "ancestries";
