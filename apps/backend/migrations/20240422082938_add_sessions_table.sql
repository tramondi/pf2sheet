-- migrate:up
CREATE TABLE "sessions" (
	"token"     char(36),
	"player_id" int not null,

	foreign key ("player_id") references "players"("id"),

	primary key ("token")
);

-- migrate:down

