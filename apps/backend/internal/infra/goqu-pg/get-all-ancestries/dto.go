package get_all_ancestries

type Ancestry struct {
	ID    int64  `db:"id"`
	Code  string `db:"code"`
	Title string `db:"title"`
}
