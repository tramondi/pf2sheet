package service

import (
	"context"
	"log/slog"

	"github.com/google/uuid"

	"github.com/alionapermes/pf2sheet/internal/domain/contract"
	"github.com/alionapermes/pf2sheet/internal/domain/entity"
)

type AuthService struct {
	logger       *slog.Logger
	playersRepo  contract.PlayersRepo
	sessionsRepo contract.SessionsRepo
}

func NewAuthService(
	logger *slog.Logger,
	playersRepo contract.PlayersRepo,
	sessionsRepo contract.SessionsRepo,
) *AuthService {
	return &AuthService{
		logger:       logger,
		playersRepo:  playersRepo,
		sessionsRepo: sessionsRepo,
	}
}

func (self *AuthService) Auth(
	ctx context.Context,
	login string,
	password string,
) (entity.Player, error) {
	player, err := self.playersRepo.FindByLogin(ctx, login)
	if err != nil {
		return entity.Player{}, err
	}

	if !player.ValidatePassword(password) {
		return entity.Player{}, contract.ErrInvalidCredentials
	}

	return player, nil
}

func (self *AuthService) CreateSession(
	ctx context.Context,
	playerID entity.PlayerID,
) (entity.Session, error) {
	token := uuid.NewString()
	session := entity.Session{PlayerID: playerID, Token: token}

	if err := self.sessionsRepo.Add(ctx, session); err != nil {
		return entity.Session{}, err
	}

	return session, nil
}
