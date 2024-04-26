package add_sheet

import "context"

type Output struct {
	SheetID int64
}

type Input struct {
	Sheet
}

func (self *db) Query(
	ctx context.Context,
	input Input,
) (Output, error) {
	insert := self.
		Insert("sheets").
		Rows(input.Sheet).
		Returning("id").
		Executor()

	var id int64

	if _, err := insert.ScanValContext(ctx, &id); err != nil {
		return Output{}, err
	}

	return Output{SheetID: id}, nil
}
