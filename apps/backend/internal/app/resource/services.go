package resource

import (
	"log/slog"

	"github.com/alionapermes/pf2sheet/internal/domain/contract"
	"github.com/alionapermes/pf2sheet/internal/domain/service"
)

type Services struct {
	Auth    contract.AuthService
	Players contract.PlayersService
}

func initServices(
	logger *slog.Logger,
	repos Repositories,
) Services {
	authService := service.NewAuthService(logger, repos.Players, repos.Sessions)
	playersService := service.NewPlayersService(logger, repos.Players)

	return Services{
		Auth:    authService,
		Players: playersService,
	}
}
