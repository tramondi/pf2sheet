package contract

import (
	"context"

	"github.com/alionapermes/pf2sheet/internal/domain/entity"
)

type PlayersRepo interface {
	Add(ctx context.Context, player entity.Player) (entity.PlayerID, error)

	GetByLogin(ctx context.Context, login string) (entity.Player, error)
}
