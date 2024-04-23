package resource

import (
	"log/slog"

	"github.com/alionapermes/pf2sheet/internal/adapter/pg"
	"github.com/alionapermes/pf2sheet/internal/domain/contract"
	"github.com/alionapermes/pf2sheet/internal/infra/sqlc-pg/dao"
)

type Repositories struct {
	Knowledge contract.KnowledgeRepo
	Sessions  contract.SessionsRepo
	Sheets    contract.SheetsRepo
	Players   contract.PlayersRepo
}

func initRepositories(
	logger *slog.Logger,
	querier dao.Querier,
) Repositories {
	knowledgeRepo := pg.NewKnowledgeRepo(logger, querier)
	sessionsRepo := pg.NewSessionsRepo(logger, querier)
	sheetsRepo := pg.NewSheetsRepo(logger, querier)
	playersRepo := pg.NewPlayersRepo(logger, querier)

	return Repositories{
		Knowledge: knowledgeRepo,
		Sessions:  sessionsRepo,
		Sheets:    sheetsRepo,
		Players:   playersRepo,
	}
}
