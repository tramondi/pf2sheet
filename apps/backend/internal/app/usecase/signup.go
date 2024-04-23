package usecase

import (
	"context"
	"log/slog"

	"github.com/alionapermes/pf2sheet/internal/domain/contract"
	"github.com/alionapermes/pf2sheet/internal/domain/entity"
)

type SignupUsecase struct {
	logger         *slog.Logger
	authService    contract.AuthService
	playersService contract.PlayersService
}

func NewSignupUsecase(
	logger *slog.Logger,
	authService contract.AuthService,
	playersService contract.PlayersService,
) SignupUsecase {
	return SignupUsecase{
		logger:         logger,
		authService:    authService,
		playersService: playersService,
	}
}

func (self *SignupUsecase) Execute(
	ctx context.Context,
	login string,
	pass string,
) (entity.Session, error) {
	player, err := self.playersService.CreatePlayer(ctx, login, pass)
	if err != nil {
		self.logger.Debug("failed to create player: %s", err)

		return entity.Session{}, err
	}

	session, err := self.authService.CreateSession(ctx, player.ID)
	if err != nil {
		self.logger.Debug("failed to create session: %s", err)

		return entity.Session{}, err
	}

	return session, nil
}
