package usecase

import (
	"context"
	"log/slog"

	"github.com/alionapermes/pf2sheet/internal/domain/contract"
	"github.com/alionapermes/pf2sheet/internal/domain/entity"
)

type GetAllSheetsUsecase struct {
	logger        *slog.Logger
	knowledgeRepo contract.KnowledgeRepo
	sheetsRepo    contract.SheetsRepo
}

func NewGetAllSheetsUsecase(
	logger *slog.Logger,
	knowledgeRepo contract.KnowledgeRepo,
	sheetsRepo contract.SheetsRepo,
) GetAllSheetsUsecase {
	return GetAllSheetsUsecase{
		logger:        logger,
		knowledgeRepo: knowledgeRepo,
		sheetsRepo:    sheetsRepo,
	}
}

func (self *GetAllSheetsUsecase) Execute(
	ctx context.Context,
	playerID entity.PlayerID,
) ([]entity.Sheet, error) {
	sheets, err := self.sheetsRepo.GetByPlayerID(ctx, playerID)
	if err != nil {
		self.logger.Error(
			"failed to get player sheets",
			slog.Int("player_id", playerID.Value()),
			slog.String("error", err.Error()),
		)
		return nil, err
	}

	return sheets, nil
}
