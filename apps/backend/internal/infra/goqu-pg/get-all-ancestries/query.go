package get_all_ancestries

import "context"

type Output struct {
	Items []struct {
		Ancestry Ancestry `db:"ancestries"`
	}
}

func (self *db) Query(ctx context.Context) (Output, error) {
	dataset := self.From("ancestries")

	var output Output

	if err := dataset.ScanStructsContext(ctx, &output.Items); err != nil {
		return Output{}, err
	}

	return output, nil
}
