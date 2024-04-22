package entity

import "golang.org/x/crypto/bcrypt"

type PlayerID int

func (self PlayerID) Value() int {
	return int(self)
}

type Player struct {
	ID          PlayerID
	DisplayName *string
	Login       string
	PassHash    string
}

func NewPlayer(login, password string) (Player, error) {
	passHashBytes, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return Player{}, err
	}

	return Player{
		Login:    login,
		PassHash: string(passHashBytes),
	}, nil
}

func (self *Player) ValidatePassword(otherPassword string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(self.PassHash),
		[]byte(otherPassword),
	)

	return err == nil
}
