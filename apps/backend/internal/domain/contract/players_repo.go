package contract

import (
	"context"

	"github.com/alionapermes/pf2sheet/internal/domain/entity"
)

type PlayersRepo interface {
	Add(ctx context.Context, player entity.Player) (entity.Player, error)

	FindByLogin(ctx context.Context, login string) (entity.Player, error)
}
