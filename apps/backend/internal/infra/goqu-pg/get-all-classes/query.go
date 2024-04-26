package get_all_classes

import "context"

type Output struct {
	Items []struct {
		Class Class `db:"classes"`
	}
}

func (self *db) Query(ctx context.Context) (Output, error) {
	dataset := self.From("classes")

	var output Output

	if err := dataset.ScanStructsContext(ctx, &output.Items); err != nil {
		return Output{}, err
	}

	return output, nil
}
