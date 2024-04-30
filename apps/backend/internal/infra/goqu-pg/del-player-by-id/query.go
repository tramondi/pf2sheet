package del_player_by_id

import (
	"context"

	"github.com/doug-martin/goqu/v9"
)

type Input struct {
	PlayerID int64
}

func (self *db) Query(ctx context.Context, input Input) error {
	_delete := self.
		Delete("player").
		Where(goqu.I("id").Eq(input.PlayerID)).
		Executor()

	if _, err := _delete.ExecContext(ctx); err != nil {
		return err
	}

	return nil
}
