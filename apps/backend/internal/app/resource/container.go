package resource

import (
	"log/slog"

	"github.com/alionapermes/pf2sheet/internal/infra/sqlc-pg/dao"
)

type Container struct {
	Repositories
	Services
	Usecases
}

func InitContainer(
	logger *slog.Logger,
	querier dao.Querier,
) (Container, error) {
	repos := initRepositories(logger, querier)
	services := initServices(logger, repos)
	usecases := initUsecases(logger, repos, services)

	return Container{
		Repositories: repos,
		Services:     services,
		Usecases:     usecases,
	}, nil
}
