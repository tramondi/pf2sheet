-- migrate:up
ALTER TABLE "sheets" ADD COLUMN "note" TEXT;

-- migrate:down

