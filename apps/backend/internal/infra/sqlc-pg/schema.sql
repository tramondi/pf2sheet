CREATE TYPE "spell_type" AS ENUM ('spell', 'focus', 'cantrip');
CREATE TYPE "feat_src" AS ENUM ('ancestry', 'skill', 'general', 'class', 'bonus');
CREATE TYPE "prof_lvl" AS ENUM ('none', 'trained', 'expert', 'master', 'legend');

CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "name" CHAR(30) NOT NULL
);

CREATE TABLE "base_blocks" (
  "id"          SERIAL PRIMARY KEY,
  "class_id"    INT    REFERENCES "classes"("id"),
  "descriptors" TEXT   NOT NULL DEFAULT '',
  "fullname"    CHAR(60),
  "ancestry"    CHAR(60),
  "background"  CHAR(60),
);

CREATE TABLE "other_blocks" (
  "id"           SERIAL PRIMARY KEY,
  "hp"           INT,
  "armor"        INT,
  "difficulty"   INT,
  "fort"         INT,
  "ref"          INT,
  "will"         INT,
  "perception"   INT,
  "senses_notes" TEXT
);

CREATE TABLE "spell_blocks" (
  "id"              SERIAL PRIMARY KEY,
  "traditions"      TEXT   NOT NULL DEFAULT '',
  "attack"          INT,
  "difficulty"      INT,
  "spells_max_"     INT,
  "spells_remains"  INT,
  "focuses_max"     INT,
  "focuses_remains" INT
);

CREATE TABLE "sheets" (
  "id"             SERIAL PRIMARY KEY,
  "user_id"        INT    NOT NULL REFERENCES "users"("id"),
  "base_block_id"  INT    NOT NULL REFERENCES "other_blocks"("id"),
  "other_block_id" INT    NOT NULL REFERENCES "other_blocks"("id"),
  "spell_block_id" INT    NOT NULL REFERENCES "spell_blocks"("id"),
  "details"        TEXT   NOT NULL DEFAULT ''
);

CREATE TABLE "skills" (
  "id"          SERIAL     PRIMARY KEY,
  "sheet_id"    INT        NOT NULL REFERENCES "sheets"("id"),
  "title"       CHAR(20)   NOT NULL,
  "lvl"         "prof_lvl" NOT NULL DEFAULT 'none',
  "item_bonus"  INT        NOT NULL DEFAULT 0,
  "armor_bonus" INT        NOT NULL DEFAULT 0
);

CREATE TABLE "abilities" (
  "id"          SERIAL   PRIMARY KEY,
  "sheet_id"    INT      NOT NULL REFERENCES "sheets"("id"),
  "title"       CHAR(20) NOT NULL,
  "modifier"    INT      NOT NULL DEFAULT 0
);

CREATE TABLE "feats" (
  "id"          SERIAL     PRIMARY KEY,
  "sheet_id"    INT        NOT NULL REFERENCES "sheets"("id"),
  "source"      "feat_src" NOT NULL,
  "title"       CHAR(60)   NOT NULL,
  "description" TEXT       NOT NULL DEFAULT ''
);

CREATE TABLE "spells" (
  "id"             SERIAL       PRIMARY KEY,
  "sheet_id"       INT          NOT NULL REFERENCES "sheets"("id"),
  "spell_block_id" INT          NOT NULL REFERENCES "spell_blocks"("id"),
  "type"           "spell_type" NOT NULL,
  "title"          CHAR(60)     NOT NULL,
  "description"    TEXT         NOT NULL DEFAULT ''
);
