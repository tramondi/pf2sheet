package usecase

import (
	"context"
	"log/slog"

	"github.com/alionapermes/pf2sheet/internal/domain/contract"
	"github.com/alionapermes/pf2sheet/internal/domain/entity"
)

type UpdatePlayerUsecase struct {
	logger      *slog.Logger
	playersRepo contract.PlayersRepo
}

func NewUpdatePlayerUsecase(
	logger *slog.Logger,
	playersRepo contract.PlayersRepo,
) UpdatePlayerUsecase {
	return UpdatePlayerUsecase{
		logger:      logger,
		playersRepo: playersRepo,
	}
}

func (self *UpdatePlayerUsecase) Execute(
	ctx context.Context,
	player entity.Player,
) error {
	if err := self.playersRepo.Update(ctx, player); err != nil {
		self.logger.Error(
			"failed to update player",
			slog.Int("player_id", player.ID.Value()),
			slog.String("error", err.Error()),
		)
		return err
	}

	return nil
}
