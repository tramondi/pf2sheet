package add_sheet

type Sheet struct {
	PlayerID   int64   `db:"player_id"`
	AncestryID *int64  `db:"ancestry_id" goqu:"omitnil"`
	ClassID    *int64  `db:"class_id" goqu:"omitnil"`
	FullName   *string `db:"fullname" goqu:"omitnil"`
	Background *string `db:"background" goqu:"omitnil"`
	Level      *int16  `db:"level" goqu:"omitnil"`
	HpCurrent  *int16  `db:"hp_current" goqu:"omitnil"`
	HpMax      *int16  `db:"hp_max" goqu:"omitnil"`
}
