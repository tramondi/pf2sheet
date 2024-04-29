package usecase

import (
	"context"
	"log/slog"

	"github.com/alionapermes/pf2sheet/internal/domain/contract"
	"github.com/alionapermes/pf2sheet/internal/domain/entity"
)

type SignoutUsecase struct {
	logger      *slog.Logger
	authService contract.AuthService
}

func NewSignoutUsecase(
	logger *slog.Logger,
	authService contract.AuthService,
) SignoutUsecase {
	return SignoutUsecase{
		logger:      logger,
		authService: authService,
	}
}

func (self *SignoutUsecase) Execute(
	ctx context.Context,
	playerID entity.PlayerID,
	token string,
) error {
	session := entity.Session{
		PlayerID: playerID,
		Token:    token,
	}

	if err := self.authService.Unauth(ctx, session); err != nil {
		self.logger.Debug("failed to create session: %s", err)

		return err
	}

	return nil
}
