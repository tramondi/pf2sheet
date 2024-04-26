package add_player

import "context"

type Input struct {
	Player
}

type Output struct {
	PlayerID int64
}

func (self *db) Query(
	ctx context.Context,
	input Input,
) (Output, error) {
	insert := self.
		Insert("players").
		Rows(input.Player).
		Returning("id").
		Executor()

	var id int64

	if _, err := insert.ScanVal(&id); err != nil {
		return Output{}, err
	}

	return Output{PlayerID: id}, nil
}
