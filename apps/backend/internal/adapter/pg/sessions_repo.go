package pg

import (
	"context"
	"log/slog"

	"github.com/doug-martin/goqu/v9"

	"github.com/alionapermes/pf2sheet/internal/domain/entity"
	add_session "github.com/alionapermes/pf2sheet/internal/infra/goqu-pg/add-session"
	delete_session_by_token "github.com/alionapermes/pf2sheet/internal/infra/goqu-pg/del-session-by-token"
	get_session_by_player_id "github.com/alionapermes/pf2sheet/internal/infra/goqu-pg/get-session-by-player-id"
	get_session_by_token "github.com/alionapermes/pf2sheet/internal/infra/goqu-pg/get-session-by-token"
)

type SessionsRepo struct {
	logger *slog.Logger
	db     *goqu.Database
}

func NewSessionsRepo(
	logger *slog.Logger,
	db *goqu.Database,
) *SessionsRepo {
	return &SessionsRepo{
		logger: logger,
		db:     db,
	}
}

func (self *SessionsRepo) Add(
	ctx context.Context,
	session entity.Session,
) error {
	input := add_session.Input{
		Session: add_session.Session{
			Token:    session.Token,
			PlayerID: int64(session.PlayerID.Value()),
		},
	}

	if err := add_session.DB(self.db).Query(ctx, input); err != nil {
		return wrapQueryError(
			err, "failed to add session for player %d", input.PlayerID)
	}

	return nil
}

func (self *SessionsRepo) GetByToken(
	ctx context.Context,
	token string,
) (entity.Session, error) {
	input := get_session_by_token.Input{Token: token}

	output, err := get_session_by_token.DB(self.db).Query(ctx, input)
	if err != nil {
		return entity.Session{},
			wrapQueryError(err, "failed to find session with token %s", token)
	}

	session := entity.Session{
		Token:    output.Session.Token,
		PlayerID: entity.PlayerID(output.Session.PlayerID),
	}

	return session, nil
}

func (self *SessionsRepo) GetByPlayerID(
	ctx context.Context,
	playerID entity.PlayerID,
) (entity.Session, error) {
	input := get_session_by_player_id.Input{PlayerID: int64(playerID)}

	output, err := get_session_by_player_id.DB(self.db).Query(ctx, input)
	if err != nil {
		return entity.Session{},
			wrapQueryError(err, "failed to find session with player id %d", playerID)
	}

	session := entity.Session{
		Token:    output.Session.Token,
		PlayerID: entity.PlayerID(output.Session.PlayerID),
	}

	return session, nil
}

func (self *SessionsRepo) DeleteByToken(
	ctx context.Context,
	token string,
) error {
	input := delete_session_by_token.Input{Token: token}

	if err := delete_session_by_token.DB(self.db).Query(ctx, input); err != nil {
		return wrapQueryError(err, "failed to find session with token %s", token)
	}

	return nil
}
