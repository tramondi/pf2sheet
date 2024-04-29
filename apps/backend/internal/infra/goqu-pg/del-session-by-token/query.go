package del_session_by_token

import (
	"context"

	"github.com/doug-martin/goqu/v9"
)

type Input struct {
	Token string
}

func (self *db) Query(ctx context.Context, input Input) error {
	_delete := self.
		Delete("sessions").
		Where(goqu.I("token").Eq(input.Token)).
		Executor()

	if _, err := _delete.ExecContext(ctx); err != nil {
		return err
	}

	return nil
}
