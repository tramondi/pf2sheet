package upd_player

import (
	"context"

	"github.com/doug-martin/goqu/v9"
)

type Input struct {
	PlayerID int64
	Player
}

func (self *db) Query(ctx context.Context, input Input) error {
	update := self.
		Update("players").
		Set(input.Player).
		Where(goqu.I("id").Eq(input.PlayerID)).
		Executor()

	if _, err := update.ExecContext(ctx); err != nil {
		return err
	}

	return nil
}
