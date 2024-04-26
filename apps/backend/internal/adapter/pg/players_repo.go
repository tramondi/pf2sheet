package pg

import (
	"context"
	"log/slog"

	"github.com/doug-martin/goqu/v9"

	"github.com/alionapermes/pf2sheet/internal/domain/entity"
	add_player "github.com/alionapermes/pf2sheet/internal/infra/goqu-pg/add-player"
	get_player_by_login "github.com/alionapermes/pf2sheet/internal/infra/goqu-pg/get-player-by-login"
)

type PlayersRepo struct {
	logger *slog.Logger
	db     *goqu.Database
}

func NewPlayersRepo(
	logger *slog.Logger,
	db *goqu.Database,
) *PlayersRepo {
	return &PlayersRepo{
		logger: logger,
		db:     db,
	}
}

func (self *PlayersRepo) Add(
	ctx context.Context,
	player entity.Player,
) (entity.PlayerID, error) {
	input := add_player.Input{
		Player: add_player.Player{
			DisplayName: player.DisplayName,
			Login:       player.Login,
			PassHash:    player.PassHash,
		},
	}

	output, err := add_player.DB(self.db).Query(ctx, input)
	if err != nil {
		return 0, wrapQueryError(err, "failed to add player")
	}

	return entity.PlayerID(output.PlayerID), nil
}

func (self *PlayersRepo) GetByLogin(
	ctx context.Context,
	login string,
) (entity.Player, error) {
	input := get_player_by_login.Input{Login: login}

	output, err := get_player_by_login.DB(self.db).Query(ctx, input)
	if err != nil {
		return entity.Player{},
			wrapQueryError(err, "failed to get player with login %s", login)
	}

	player := entity.Player{
		ID:       entity.PlayerID(output.ID),
		Login:    output.Login,
		PassHash: output.PassHash,
	}

	return player, nil
}
