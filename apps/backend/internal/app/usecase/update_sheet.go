package usecase

import (
	"context"
	"errors"
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
	_sheet, err := self.sheetsRepo.GetByID(ctx, sheet.ID)
	if err != nil {
		self.logger.Error(
			"failed to get sheet by id",
			slog.Int("sheet_id", int(sheet.ID)),
			slog.String("error", err.Error()),
		)
		return err
	}

	if _sheet.PlayerID != sheet.PlayerID {
		self.logger.Error(
			"sheet has unexpected player id",
			slog.Int("input.player_id", sheet.PlayerID.Value()),
			slog.Int("actual.player_id", _sheet.PlayerID.Value()),
		)
		return errors.New("unexpected player id")
	}

	if err := self.sheetsRepo.Update(ctx, sheet); err != nil {
		self.logger.Error(
			"failed to update sheet", slog.Int("sheet_id", sheet.ID.Value()))
		return err
	}

	return nil
}
