package contract

import (
	"context"

	"github.com/alionapermes/pf2sheet/internal/domain/entity"
)

type SessionsRepo interface {
	Add(ctx context.Context, session entity.Session) error

	GetByToken(ctx context.Context, token string) (entity.Session, error)

	GetByPlayerID(
		ctx context.Context,
		playerID entity.PlayerID,
	) (entity.Session, error)

	DeleteByToken(ctx context.Context, token string) error
}
