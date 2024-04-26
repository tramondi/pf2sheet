package resource

import (
	"log/slog"

	"github.com/doug-martin/goqu/v9"

	"github.com/alionapermes/pf2sheet/internal/adapter/pg"
	"github.com/alionapermes/pf2sheet/internal/domain/contract"
)

type Repositories struct {
	Knowledge contract.KnowledgeRepo
	Sessions  contract.SessionsRepo
	Sheets    contract.SheetsRepo
	Players   contract.PlayersRepo
}

func initRepositories(
	logger *slog.Logger,
	goquDB *goqu.Database,
) Repositories {
	knowledgeRepo := pg.NewKnowledgeRepo(logger, goquDB)
	sessionsRepo := pg.NewSessionsRepo(logger, goquDB)
	sheetsRepo := pg.NewSheetsRepo(logger, goquDB)
	playersRepo := pg.NewPlayersRepo(logger, goquDB)

	return Repositories{
		Knowledge: knowledgeRepo,
		Sessions:  sessionsRepo,
		Sheets:    sheetsRepo,
		Players:   playersRepo,
	}
}
