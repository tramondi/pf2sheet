package get_sheet_by_id

import (
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
)

type Output struct {
	Sheet    Sheet    `db:"sheets"`
	Ancestry Ancestry `db:"ancestries"`
	Class    Class    `db:"classes"`
}

type Input struct {
	SheetID int64
}

func (self *db) Query(
	ctx context.Context,
	input Input,
) (Output, error) {
	dataset := self.
		From("sheets").
		LeftJoin(
			goqu.T("ancestries"),
			goqu.On(goqu.I("sheets.ancestry_id").Eq(goqu.I("ancestries.id"))),
		).
		LeftJoin(
			goqu.T("classes"),
			goqu.On(goqu.I("sheets.class_id").Eq(goqu.I("classes.id"))),
		).
		Where(goqu.I("sheets.id").Eq(input.SheetID))

	var output Output

	found, err := dataset.ScanStructContext(ctx, &output)
	if err != nil {
		return Output{}, err
	}

	if !found {
		return Output{}, sql.ErrNoRows
	}

	return output, nil
}
