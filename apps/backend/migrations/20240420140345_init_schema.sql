-- migrate:up
DROP TABLE IF EXISTS "players" cascade;
DROP TABLE IF EXISTS "ancestries" cascade;
DROP TABLE IF EXISTS "classes" cascade;
DROP TABLE IF EXISTS "sheets" cascade;

CREATE TABLE "players" (
  "id"        serial,
  "name"      char(30) not null,
  "pass_hash" char(72) not null,

  primary key ("id")
);

CREATE TABLE "ancestries" (
  "id"    serial,
  "code"  char(30) unique not null,
  "title" char(30)        not null,

  primary key ("id")
);
CREATE INDEX "idx_ancestries_code" ON "ancestries" USING HASH ("code");

CREATE TABLE "classes" (
  "id"    serial,
  "code"  char(30) unique not null,
  "title" char(30)        not null,

  primary key ("id")
);
CREATE INDEX "idx_classes_code" ON "classes" USING HASH ("code");

CREATE TABLE "sheets" (
  "id"          serial,
  "player_id"   int      not null,
  "ancestry_id" int,
  "class_id"    int,
  "background"  char(60),
  "fullname"    char(60),
  "level"       smallint,
  "hp_current"  smallint,
  "hp_max"      smallint,

  foreign key ("player_id")   references "players"("id")
    on delete cascade,
  foreign key ("ancestry_id") references "ancestries"("id")
    on delete set null,
  foreign key ("class_id")    references "classes"("id")
    on delete set null,

  primary key ("id")
);

-- migrate:down

