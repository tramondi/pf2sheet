package get_session_by_player_id

import (
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
)

type Input struct {
	PlayerID int64
}

type Output struct {
	Session
}

func (self *db) Query(
	ctx context.Context,
	input Input,
) (Output, error) {
	dataset := self.
		From("sessions").
		Where(goqu.I("player_id").Eq(input.PlayerID))

	var session Session

	found, err := dataset.ScanStructContext(ctx, &session)
	if err != nil {
		return Output{}, err
	}

	if !found {
		return Output{}, sql.ErrNoRows
	}

	return Output{Session: session}, nil
}
