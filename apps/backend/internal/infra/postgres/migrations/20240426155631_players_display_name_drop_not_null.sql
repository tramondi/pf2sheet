-- migrate:up
ALTER TABLE "players" ALTER COLUMN "display_name" DROP NOT NULL;

-- migrate:down

