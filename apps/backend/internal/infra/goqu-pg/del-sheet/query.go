package del_sheet

import (
	"context"

	"github.com/doug-martin/goqu/v9"
)

type Input struct {
	SheetID int64
}

func (self *db) Query(ctx context.Context, input Input) error {
	_delete := self.
		Delete("sheets").
		Where(goqu.I("id").Eq(input.SheetID)).
		Executor()

	if _, err := _delete.ExecContext(ctx); err != nil {
		return err
	}

	return nil
}
