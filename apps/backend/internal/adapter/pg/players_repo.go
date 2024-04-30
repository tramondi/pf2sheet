package pg

import (
	"context"
	"log/slog"

	"github.com/doug-martin/goqu/v9"

	"github.com/alionapermes/pf2sheet/internal/domain/entity"
	add_player "github.com/alionapermes/pf2sheet/internal/infra/goqu-pg/add-player"
	del_player_by_id "github.com/alionapermes/pf2sheet/internal/infra/goqu-pg/del-player-by-id"
	get_player_by_id "github.com/alionapermes/pf2sheet/internal/infra/goqu-pg/get-player-by-id"
	get_player_by_login "github.com/alionapermes/pf2sheet/internal/infra/goqu-pg/get-player-by-login"
	upd_player "github.com/alionapermes/pf2sheet/internal/infra/goqu-pg/upd-player"
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

func (self *PlayersRepo) DeleteByID(
	ctx context.Context,
	playerID entity.PlayerID,
) error {
	input := del_player_by_id.Input{PlayerID: int64(playerID)}

	if err := del_player_by_id.DB(self.db).Query(ctx, input); err != nil {
		return wrapQueryError(err, "failed to delete player with id %d", playerID)
	}

	return nil
}

func (self *PlayersRepo) Update(
	ctx context.Context,
	player entity.Player,
) error {
	input := upd_player.Input{
		PlayerID: int64(player.ID),
		Player: upd_player.Player{
			DisplayName: player.DisplayName,
		},
	}

	if err := upd_player.DB(self.db).Query(ctx, input); err != nil {
		return wrapQueryError(err, "failed to update player with id %d", player.ID)
	}

	return nil
}

func (self *PlayersRepo) GetByID(
	ctx context.Context,
	playerID entity.PlayerID,
) (entity.Player, error) {
	input := get_player_by_id.Input{ID: int64(playerID)}

	output, err := get_player_by_id.DB(self.db).Query(ctx, input)
	if err != nil {
		return entity.Player{}, err
	}

	player := entity.Player{
		ID:          entity.PlayerID(output.ID),
		DisplayName: output.DisplayName,
		Login:       output.Login,
		PassHash:    output.PassHash,
	}

	return player, nil
}
