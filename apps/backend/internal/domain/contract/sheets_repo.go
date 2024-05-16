package contract

import (
	"context"

	"github.com/alionapermes/pf2sheet/internal/domain/entity"
)

type SheetsRepo interface {
	GetByPlayerID(
		ctx context.Context,
		playerID entity.PlayerID,
	) ([]entity.Sheet, error)

	Add(
		ctx context.Context,
		playerID entity.PlayerID,
		sheet entity.Sheet,
	) (entity.SheetID, error)

	DeleteByID(ctx context.Context, id entity.SheetID) error

	Update(ctx context.Context, sheet entity.Sheet) error

	GetByID(ctx context.Context, sheetID entity.SheetID) (entity.Sheet, error)
}
