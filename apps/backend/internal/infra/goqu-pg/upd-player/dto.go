package upd_player

type Player struct {
	DisplayName *string `db:"display_name" goqu:"omitnil"`
}
