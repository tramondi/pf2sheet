package upd_sheet

import (
	"context"

	"github.com/doug-martin/goqu/v9"
)

type Input struct {
	SheetID int64
	Sheet
}

func (self *db) Query(ctx context.Context, input Input) error {
	update := self.
		Update("sheets").
		Set(input.Sheet).
		Where(goqu.I("id").Eq(input.SheetID)).
		Executor()

	if _, err := update.ExecContext(ctx); err != nil {
		return err
	}

	return nil
}
