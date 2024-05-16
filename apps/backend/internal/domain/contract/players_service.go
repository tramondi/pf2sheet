package contract

import (
	"context"

	"github.com/alionapermes/pf2sheet/internal/domain/entity"
)

type PlayersService interface {
	CreatePlayer(
		ctx context.Context,
		login string,
		password string,
		displayName *string,
	) (entity.Player, error)
}
