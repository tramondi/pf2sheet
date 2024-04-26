package add_player

type Player struct {
	DisplayName *string `db:"display_name" goqu:"omitnil"`
	PassHash    string  `db:"pass_hash"`
	Login       string  `db:"login"`
}
