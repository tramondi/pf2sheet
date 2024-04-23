package usecase

import (
	"context"
	"errors"
	"log/slog"

	"github.com/alionapermes/pf2sheet/internal/domain/contract"
	"github.com/alionapermes/pf2sheet/internal/domain/entity"
)

type SigninUsecase struct {
	logger      *slog.Logger
	authService contract.AuthService
	sheetsRepo  contract.SheetsRepo
}

func NewSigninUsecase(
	logger *slog.Logger,
	authService contract.AuthService,
	sheetsRepo contract.SheetsRepo,
) SigninUsecase {
	return SigninUsecase{
		logger:      logger,
		authService: authService,
		sheetsRepo:  sheetsRepo,
	}
}

func (self *SigninUsecase) Execute(
	ctx context.Context,
	login string,
	password string,
) (entity.Session, error) {
	player, err := self.authService.Auth(ctx, login, password)
	if err != nil {
		if !errors.Is(err, contract.ErrInvalidCredentials) {
			self.logger.Error(
				"failed to auth player",
				slog.String("login", player.Login),
				slog.String("error", err.Error()),
			)
		}

		return entity.Session{}, err
	}

	session, err := self.authService.GetOrCreateSession(ctx, player.ID)
	if err != nil {
		return entity.Session{}, err
	}

	// sheets, err := self.sheetsRepo.GetByPlayerID(ctx, player.ID)
	// if err != nil {
	// 	if !errors.Is(err, contract.ErrNotFound) {
	// 		self.logger.Error(
	// 			"failed to get sheets by player id",
	// 			slog.Int("player_id", player.ID.Value()),
	// 			slog.String("error", err.Error()),
	// 		)
	// 	}
	//
	// 	return session, player, nil, nil
	// }

	return session, nil
}
