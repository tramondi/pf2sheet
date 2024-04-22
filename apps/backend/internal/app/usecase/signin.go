package usecase

import (
	"context"
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
) (entity.Session, entity.Player, []entity.Sheet, error) {
	player, err := self.authService.Auth(ctx, login, password)
	if err != nil {
		return entity.Session{}, entity.Player{}, nil, err
	}

	session, err := self.authService.CreateSession(ctx, player.ID)
	if err != nil {
		return entity.Session{}, entity.Player{}, nil, err
	}

	sheets, err := self.sheetsRepo.GetByPlayerID(ctx, player.ID)
	if err != nil {
		// @todo: self.logger.Errorf(…)
	}

	return session, player, sheets, nil
}
