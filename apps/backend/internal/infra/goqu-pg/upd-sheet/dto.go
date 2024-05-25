package upd_sheet

type Sheet struct {
	AncestryID *int64  `db:"ancestry_id"`
	ClassID    *int64  `db:"class_id"`
	FullName   *string `db:"fullname"`
	Background *string `db:"background"`
	Level      *int16  `db:"level"`
	HpCurrent  *int16  `db:"hp_current"`
	HpMax      *int16  `db:"hp_max"`
	Note       *string `db:"note"`
}
