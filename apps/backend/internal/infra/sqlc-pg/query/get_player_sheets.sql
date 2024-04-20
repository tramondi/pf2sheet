-- name: GetPlayerSheets :many
SELECT
  sqlc.embed(sheets),
  sqlc.embed(players),
  sqlc.embed(ancestries),
  sqlc.embed(classes)
FROM "sheets"
  LEFT JOIN "players"
    ON "sheets"."player_id" = "players"."id"
  LEFT JOIN "ancestries"
    ON "sheets"."ancestry_id" = "ancestries"."id"
  LEFT JOIN "classes"
    ON "sheets"."class_id" = "classes"."id"
WHERE "sheets"."id" = $1;
