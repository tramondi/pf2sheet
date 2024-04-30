package contract

import (
	"context"

	"github.com/alionapermes/pf2sheet/internal/domain/entity"
)

type PlayersRepo interface {
	Add(ctx context.Context, player entity.Player) (entity.PlayerID, error)

	GetByLogin(ctx context.Context, login string) (entity.Player, error)

	GetByID(ctx context.Context, playerID entity.PlayerID) (entity.Player, error)

	Update(ctx context.Context, player entity.Player) error

	DeleteByID(ctx context.Context, playerID entity.PlayerID) error
}
