-- migrate:up
ALTER TABLE "players" RENAME COLUMN "name" TO "display_name";

ALTER TABLE "players" ADD COLUMN "login" char(30) not null;
CREATE INDEX "idx_players_login" ON "players" USING HASH ("login");

-- migrate:down

