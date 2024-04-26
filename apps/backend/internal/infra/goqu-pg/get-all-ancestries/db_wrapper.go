package get_all_ancestries

import "github.com/doug-martin/goqu/v9"

type db struct{ *goqu.Database }

func DB(database *goqu.Database) *db { return &db{Database: database} }
