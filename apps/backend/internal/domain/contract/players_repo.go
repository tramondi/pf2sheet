package contract

import "context"

type PlayersRepo interface {
	GetPlayerByID(ctx context.Context, id int) (interface{}, error)

	AddPlayer(ctx context.Context, player interface{}) (int, error)
}
