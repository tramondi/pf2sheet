package pg

import (
	"context"
	"log/slog"

	"github.com/AlekSi/pointer"

	"github.com/alionapermes/pf2sheet/internal/domain/entity"
	"github.com/alionapermes/pf2sheet/internal/infra/sqlc-pg/dao"
)

type PlayersRepo struct {
	logger  *slog.Logger
	querier dao.Querier
}

func NewPlayersRepo(
	logger *slog.Logger,
	querier dao.Querier,
) *PlayersRepo {
	return &PlayersRepo{
		logger:  logger,
		querier: querier,
	}
}

func (self *PlayersRepo) Add(
	ctx context.Context,
	player entity.Player,
) (entity.Player, error) {
	var displayName string
	if player.DisplayName != nil {
		displayName = *player.DisplayName
	}

	params := dao.AddPlayerParams{
		DisplayName: displayName,
		PassHash:    player.PassHash,
		Login:       player.Login,
	}

	if err := self.querier.AddPlayer(ctx, params); err != nil {
		return entity.Player{}, nil
	}

	// @todo sqlc хуйня ебаная
	row, err := self.querier.FindPlayerByLogin(ctx, player.Login)
	if err != nil {
		return entity.Player{}, nil
	}

	playerOut := entity.Player{
		ID:          entity.PlayerID(row.Player.ID),
		DisplayName: pointer.ToStringOrNil(row.Player.DisplayName),
		PassHash:    row.Player.PassHash,
		Login:       row.Player.Login,
	}

	return playerOut, nil
}

func (self *PlayersRepo) FindByLogin(
	ctx context.Context,
	login string,
) (entity.Player, error) {
	row, err := self.querier.FindPlayerByLogin(ctx, login)
	if err != nil {
		return entity.Player{},
			wrapQueryError(err, "failed to find player with login %s", login)
	}

	var displayName *string
	if name := row.Player.DisplayName; name != "" {
		displayName = &name
	}

	player := entity.Player{
		ID:          entity.PlayerID(row.Player.ID),
		Login:       row.Player.Login,
		PassHash:    row.Player.PassHash,
		DisplayName: displayName,
	}

	return player, nil
}
