package get_sheets_by_player_id

type Ancestry struct {
	ID    *int64  `db:"id"`
	Code  *string `db:"code"`
	Title *string `db:"title"`
}

type Class struct {
	ID    *int64  `db:"id"`
	Code  *string `db:"code"`
	Title *string `db:"title"`
}

type Sheet struct {
	ID         int64   `db:"id"`
	PlayerID   int64   `db:"player_id"`
	AncestryID *int64  `db:"ancestry_id"`
	ClassID    *int64  `db:"class_id"`
	FullName   *string `db:"fullname"`
	Background *string `db:"background"`
	Level      *int16  `db:"level"`
	HpCurrent  *int16  `db:"hp_current"`
	HpMax      *int16  `db:"hp_max"`
	Note       *string `db:"note"`
}
