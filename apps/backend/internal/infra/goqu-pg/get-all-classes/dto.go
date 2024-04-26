package get_all_classes

type Class struct {
	ID    int64  `db:"id"`
	Code  string `db:"code"`
	Title string `db:"title"`
}
