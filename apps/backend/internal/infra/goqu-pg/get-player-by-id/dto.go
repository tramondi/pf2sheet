package get_player_by_id

type Player struct {
	ID          int64   `db:"id"`
	Login       string  `db:"login"`
	PassHash    string  `db:"pass_hash"`
	DisplayName *string `db:"display_name"`
}
