package usecase

import (
	"context"
	"log/slog"

	"github.com/alionapermes/pf2sheet/internal/domain/contract"
	"github.com/alionapermes/pf2sheet/internal/domain/entity"
)

type CreateSheetUsecase struct {
	logger        *slog.Logger
	sheetsRepo    contract.SheetsRepo
	knowledgeRepo contract.KnowledgeRepo
}

func NewCreateSheetsUsecase(
	logger *slog.Logger,
	sheetsRepo contract.SheetsRepo,
	knowledgeRepo contract.KnowledgeRepo,
) CreateSheetUsecase {
	return CreateSheetUsecase{
		logger:        logger,
		sheetsRepo:    sheetsRepo,
		knowledgeRepo: knowledgeRepo,
	}
}

func (self *CreateSheetUsecase) Execute(
	ctx context.Context,
	playerID entity.PlayerID,
	sheet entity.Sheet,
) (entity.SheetID, error) {
	sheetID, err := self.sheetsRepo.Add(ctx, playerID, sheet)
	if err != nil {
		self.logger.Error("failed to add sheet", slog.String("error", err.Error()))
		return 0, err
	}

	return sheetID, nil
}
