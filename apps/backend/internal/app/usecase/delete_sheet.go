package usecase

import (
	"context"
	"log/slog"

	"github.com/alionapermes/pf2sheet/internal/domain/contract"
	"github.com/alionapermes/pf2sheet/internal/domain/entity"
)

type DeleteSheetUsecase struct {
	logger     *slog.Logger
	sheetsRepo contract.SheetsRepo
}

func NewDeleteSheetsUsecase(
	logger *slog.Logger,
	sheetsRepo contract.SheetsRepo,
) DeleteSheetUsecase {
	return DeleteSheetUsecase{
		logger:     logger,
		sheetsRepo: sheetsRepo,
	}
}

func (self *DeleteSheetUsecase) Execute(
	ctx context.Context,
	playerID entity.PlayerID,
	sheetID entity.SheetID,
) error {
	if err := self.sheetsRepo.DeleteByID(ctx, sheetID); err != nil {
		self.logger.Error("failed to add sheet", slog.String("error", err.Error()))
		return err
	}

	return nil
}
