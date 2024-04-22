-- migrate:up
ALTER TABLE "sessions" ALTER COLUMN "token" TYPE char(72);

-- migrate:down

