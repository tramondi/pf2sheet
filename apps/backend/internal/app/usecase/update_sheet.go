package usecase

import (
	"context"
	"log/slog"

	"github.com/alionapermes/pf2sheet/internal/domain/contract"
	"github.com/alionapermes/pf2sheet/internal/domain/entity"
)

type UpdateSheetUsecase struct {
	logger     *slog.Logger
	sheetsRepo contract.SheetsRepo
}

func NewUpdateSheetUsecase(
	logger *slog.Logger,
	sheetsRepo contract.SheetsRepo,
) UpdateSheetUsecase {
	return UpdateSheetUsecase{
		logger:     logger,
		sheetsRepo: sheetsRepo,
	}
}

func (self *UpdateSheetUsecase) Execute(
	ctx context.Context,
	sheet entity.Sheet,
) error {
	if err := self.sheetsRepo.Update(ctx, sheet); err != nil {
		self.logger.Error(
			"failed to update sheet", slog.Int("sheet_id", sheet.ID.Value()))
		return err
	}

	return nil
}
