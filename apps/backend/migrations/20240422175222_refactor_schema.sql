-- migrate:up
ALTER TABLE "sessions"
	DROP CONSTRAINT "sessions_player_id_fkey",
	ADD CONSTRAINT "sessions_player_id_fkey"
		foreign key ("player_id") references "players"("id")
			on delete cascade;

DELETE FROM "players" CASCADE;

ALTER TABLE "players" ADD CONSTRAINT "players_login_key" UNIQUE ("login");

ALTER TABLE "players"
	ALTER COLUMN "display_name" TYPE varchar(30),
	ALTER COLUMN "login"        TYPE varchar(30),
	ALTER COLUMN "pass_hash"    TYPE varchar(72);

ALTER TABLE "ancestries"
	ALTER COLUMN "code"  TYPE varchar(30),
	ALTER COLUMN "title" TYPE varchar(30);

ALTER TABLE "classes"
	ALTER COLUMN "code"  TYPE varchar(30),
	ALTER COLUMN "title" TYPE varchar(30);

ALTER TABLE "sheets"
	ALTER COLUMN "background"  TYPE varchar(60),
	ALTER COLUMN "fullname"    TYPE varchar(60);

ALTER TABLE "sessions" ALTER COLUMN "token" TYPE varchar(72);

-- migrate:down

