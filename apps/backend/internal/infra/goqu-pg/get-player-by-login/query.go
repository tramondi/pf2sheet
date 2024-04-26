package get_player_by_login

import (
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
)

type Output struct {
	Player
}

type Input struct {
	Login string
}

func (self *db) Query(
	ctx context.Context,
	input Input,
) (Output, error) {
	dataset := self.
		From("players").
		Where(goqu.I("login").Eq(input.Login))

	var output Output

	found, err := dataset.ScanStructContext(ctx, &output.Player)
	if err != nil {
		return Output{}, err
	}

	if !found {
		return Output{}, sql.ErrNoRows
	}

	return output, nil
}
