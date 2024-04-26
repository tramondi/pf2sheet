package resource

import (
	"log/slog"

	"github.com/doug-martin/goqu/v9"
)

type Container struct {
	Repositories
	Services
	Usecases
}

func InitContainer(
	logger *slog.Logger,
	goquDB *goqu.Database,
) (Container, error) {
	repos := initRepositories(logger, goquDB)
	services := initServices(logger, repos)
	usecases := initUsecases(logger, repos, services)

	return Container{
		Repositories: repos,
		Services:     services,
		Usecases:     usecases,
	}, nil
}
