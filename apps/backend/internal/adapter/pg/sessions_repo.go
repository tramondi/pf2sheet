package pg

import (
	"context"
	"log/slog"

	"github.com/alionapermes/pf2sheet/internal/domain/entity"
	"github.com/alionapermes/pf2sheet/internal/infra/sqlc-pg/dao"
)

type SessionsRepo struct {
	logger  *slog.Logger
	querier dao.Querier
}

func NewSessionsRepo(
	logger *slog.Logger,
	querier dao.Querier,
) *SessionsRepo {
	return &SessionsRepo{
		logger:  logger,
		querier: querier,
	}
}

func (self *SessionsRepo) Add(
	ctx context.Context,
	session entity.Session,
) error {
	params := dao.AddSessionParams{
		Token:    session.Token,
		PlayerID: int32(session.PlayerID),
	}

	if err := self.querier.AddSession(ctx, params); err != nil {
		return err
	}

	return nil
}

func (self *SessionsRepo) FindByToken(
	ctx context.Context,
	token string,
) (entity.Session, error) {
	row, err := self.querier.FindSessionByToken(ctx, token)
	if err != nil {
		return entity.Session{},
			wrapQueryError(err, "failed to find session with token %s", token)
	}

	session := entity.Session{
		PlayerID: entity.PlayerID(row.Session.PlayerID),
		Token:    row.Session.Token,
	}

	return session, nil
}

func (self *SessionsRepo) DeleteByToken(
	ctx context.Context,
	token string,
) error {
	if err := self.querier.DeleteSessionByToken(ctx, token); err != nil {
		return wrapQueryError(err, "failed to find session with token %s", token)
	}

	return nil
}
