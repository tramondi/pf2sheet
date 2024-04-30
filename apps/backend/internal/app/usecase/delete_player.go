package usecase

import (
	"context"
	"log/slog"

	"github.com/alionapermes/pf2sheet/internal/domain/contract"
	"github.com/alionapermes/pf2sheet/internal/domain/entity"
)

type DeletePlayerUsecase struct {
	logger      *slog.Logger
	playersRepo contract.PlayersRepo
}

func NewDeletePlayerUsecase(
	logger *slog.Logger,
	playersRepo contract.PlayersRepo,
) DeletePlayerUsecase {
	return DeletePlayerUsecase{
		logger:      logger,
		playersRepo: playersRepo,
	}
}

func (self *DeletePlayerUsecase) Execute(
	ctx context.Context,
	playerID entity.PlayerID,
) error {
	if err := self.playersRepo.DeleteByID(ctx, playerID); err != nil {
		self.logger.Error(
			"failed to delete player",
			slog.Int("player_id", playerID.Value()),
			slog.String("error", err.Error()),
		)
		return err
	}

	return nil
}
