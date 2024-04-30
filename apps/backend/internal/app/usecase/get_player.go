package usecase

import (
	"context"
	"log/slog"

	"github.com/alionapermes/pf2sheet/internal/domain/contract"
	"github.com/alionapermes/pf2sheet/internal/domain/entity"
)

type GetPlayerUsecase struct {
	logger      *slog.Logger
	playersRepo contract.PlayersRepo
}

func NewGetPlayerUsecase(
	logger *slog.Logger,
	playersRepo contract.PlayersRepo,
) GetPlayerUsecase {
	return GetPlayerUsecase{
		logger:      logger,
		playersRepo: playersRepo,
	}
}

func (self *GetPlayerUsecase) Execute(
	ctx context.Context,
	playerID entity.PlayerID,
) (entity.Player, error) {
	player, err := self.playersRepo.GetByID(ctx, playerID)
	if err != nil {
		self.logger.Error(
			"failed to get player by id", slog.Int("player_id", playerID.Value()))
		return entity.Player{}, err
	}

	return player, nil
}
