package add_session

type Session struct {
	Token    string `db:"token"`
	PlayerID int64  `db:"player_id"`
}
