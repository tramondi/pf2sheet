package get_sheets_by_player_id

import (
	"context"

	"github.com/doug-martin/goqu/v9"
)

type Output struct {
	Items []struct {
		Sheet    Sheet    `db:"sheets"`
		Ancestry Ancestry `db:"ancestries"`
		Class    Class    `db:"classes"`
	}
}

type Input struct {
	PlayerID int64
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
		Where(goqu.I("sheets.player_id").Eq(input.PlayerID))

	var output Output

	if err := dataset.ScanStructsContext(ctx, &output.Items); err != nil {
		return Output{}, err
	}

	return output, nil
}
