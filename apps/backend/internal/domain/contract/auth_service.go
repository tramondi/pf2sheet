package contract

import (
	"context"

	"github.com/alionapermes/pf2sheet/internal/domain/entity"
)

type AuthService interface {
	Auth(ctx context.Context, login, password string) (entity.Player, error)

	CreateSession(ctx context.Context, playerID entity.PlayerID) (entity.Session, error)
}
