package add_session

import "context"

type Input struct {
	Session
}

func (self *db) Query(
	ctx context.Context,
	input Input,
) error {
	insert := self.
		Insert("sessions").
		Rows(input.Session).
		Executor()

	if _, err := insert.ExecContext(ctx); err != nil {
		return err
	}

	return nil
}
