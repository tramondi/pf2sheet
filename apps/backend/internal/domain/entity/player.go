package entity

type PlayerID int

func (self PlayerID) Value() int {
	return int(self)
}

type Player struct {
	ID       PlayerID
	Name     string
	PassHash string

	Sheets []Sheet
}
