package get_session_by_token

type Session struct {
	Token    string `db:"token"`
	PlayerID int64  `db:"player_id"`
}
