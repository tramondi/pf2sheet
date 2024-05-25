package entity

type SheetID int

func (self SheetID) Value() int {
	return int(self)
}

type Sheet struct {
	ID         SheetID
	PlayerID   PlayerID
	Ancestry   *Ancestry
	Class      *Class
	Background *string
	FullName   *string
	Level      *int16
	HpCurrent  *int16
	HpMax      *int16
	Note       *string
}
