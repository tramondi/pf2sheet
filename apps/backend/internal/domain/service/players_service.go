package service

import (
	"context"
	"log/slog"

	"github.com/alionapermes/pf2sheet/internal/domain/contract"
	"github.com/alionapermes/pf2sheet/internal/domain/entity"
)

type PlayersService struct {
	logger      *slog.Logger
	playersRepo contract.PlayersRepo
}

func NewPlayersService(
	logger *slog.Logger,
	playersRepo contract.PlayersRepo,
) *PlayersService {
	return &PlayersService{
		logger:      logger,
		playersRepo: playersRepo,
	}
}

func (self *PlayersService) CreatePlayer(
	ctx context.Context,
	login string,
	password string,
) (entity.Player, error) {
	_, err := self.playersRepo.GetByLogin(ctx, login)
	if err == nil {
		return entity.Player{}, contract.ErrAlreadyExists
	}

	player, err := entity.NewPlayer(login, password)
	if err != nil {
		return entity.Player{}, err
	}

	playerID, err := self.playersRepo.Add(ctx, player)
	if err != nil {
		return entity.Player{}, err
	}

	player.ID = playerID

	return player, nil
}
